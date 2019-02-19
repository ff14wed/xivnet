package xivnet

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"time"
)

// Decoder implements an FFXIV frame decoder. It is not thread-safe, so consumers
// need to be sure not to use it concurrently.
type Decoder struct {
	bufReader *bufio.Reader
	blockBuf  []byte
}

// NewDecoder creates a new instance of a frame decoder.
// bufSize controls the total size of the buffer used to store a single
// frame. It's recommended to keep this value at least 8192, since this value
// generally works for all frames (as far as I can tell).
func NewDecoder(r io.Reader, bufSize int) *Decoder {
	return &Decoder{
		bufReader: bufio.NewReaderSize(r, bufSize),

		// This buffer should be at least big enough to store all uncompressed
		// blocks
		blockBuf: make([]byte, bufSize),
	}
}

// CheckHeader checks to see whether or not the data in the buffer has a
// valid header
func (d *Decoder) CheckHeader() ([]byte, error) {
	if 28 > d.bufReader.Size() {
		return nil, InvalidFrameLengthError{length: 28, maxLength: d.bufReader.Size()}
	}
	// Validation that the frame at least has
	header, err := d.bufReader.Peek(28)
	if err != nil {
		return nil, EOFError{
			operation:       "peeking header",
			attemptedLength: 28,
			wrapped:         err,
		}
	}
	if !isValidHeader(header) {
		return nil, InvalidHeaderError{
			header: hex.EncodeToString(header),
		}
	}
	return header, nil
}

// NextFrame reads data from the underlying buffer and returns a single decoded
// FFXIV frame from the provided buffer.
func (d *Decoder) NextFrame() (*Frame, error) {
	header, err := d.CheckHeader()
	if err != nil {
		return nil, err
	}
	length := binary.LittleEndian.Uint32(header[24:])
	if length > uint32(d.bufReader.Size()) {
		return nil, InvalidFrameLengthError{length: length, maxLength: d.bufReader.Size()}
	}
	intLength := int(length)
	frameBytes, err := d.bufReader.Peek(intLength)
	if err != nil {
		return nil, EOFError{
			operation:       "peeking data",
			attemptedLength: intLength,
			wrapped:         err,
		}
	}
	defer func() {
		_, _ = d.bufReader.Discard(intLength)
	}()
	f, err := decodeFrame(frameBytes, d.blockBuf, length)
	if err != nil {
		return nil, DecodingError{
			wrapped:   err,
			debugData: hex.EncodeToString(frameBytes),
		}
	}
	return f, nil
}

func decodeFrame(frameBytes []byte, blockBuffer []byte, length uint32) (*Frame, error) {
	// Build the frame
	frame := &Frame{}
	copy(frame.Preamble[:], frameBytes[0:16])
	msecSinceEpoch := time.Duration(binary.LittleEndian.Uint64(frameBytes[16:24])) * time.Millisecond
	frame.Time = time.Unix(0, 0).Add(msecSinceEpoch)
	frame.Length = length
	frame.ConnectionType = binary.LittleEndian.Uint16(frameBytes[28:30])
	frame.Count = binary.LittleEndian.Uint16(frameBytes[30:32])
	frame.Reserved1 = frameBytes[32]
	frame.Compression = frameBytes[33]
	frame.Reserved2 = binary.LittleEndian.Uint32(frameBytes[34:38])
	frame.Reserved3 = binary.LittleEndian.Uint16(frameBytes[38:40])

	blockData := frameBytes[40:length]
	if frame.Compression > 0 {
		r, err := zlib.NewReader(bytes.NewReader(blockData))
		if err != nil {
			return nil, fmt.Errorf("error decompressing data: %s", err.Error())
		}
		n, err := r.Read(blockBuffer)
		r.Close()
		if err != io.EOF {
			return nil, fmt.Errorf("error reading decompressed data: %s", err.Error())
		}
		blockData = blockBuffer[:n]
	}

	if len(blockData) == 0 {
		return frame, nil
	}

	for {
		block, err := decodeBlock(blockData)
		if err != nil {
			return nil, fmt.Errorf("error decoding blocks: %s", err.Error())
		}
		frame.Blocks = append(frame.Blocks, block)
		blockData = blockData[block.Length:]
		if len(blockData) == 0 {
			break
		}
	}
	return frame, nil
}

func decodeBlock(blocksBytes []byte) (*Block, error) {
	cap := len(blocksBytes)
	if cap < 4 {
		return nil, errors.New("missing block data")
	}
	length := binary.LittleEndian.Uint32(blocksBytes[:4])
	if int(length) > cap {
		return nil, fmt.Errorf("not enough data: expected %d bytes, got %d", length, cap)
	}
	block := &Block{}
	block.Length = length
	block.SubjectID = binary.LittleEndian.Uint32(blocksBytes[4:8])
	block.CurrentID = binary.LittleEndian.Uint32(blocksBytes[8:12])
	block.Type = binary.LittleEndian.Uint16(blocksBytes[12:14])
	block.Pad1 = binary.LittleEndian.Uint16(blocksBytes[14:16])
	var blockData GenericBlockData
	if block.Type == BlockTypeIPC {
		block.Reserved = binary.LittleEndian.Uint16(blocksBytes[16:18])
		block.Opcode = binary.LittleEndian.Uint16(blocksBytes[18:20])
		block.Pad2 = binary.LittleEndian.Uint16(blocksBytes[20:22])
		block.ServerID = binary.LittleEndian.Uint16(blocksBytes[22:24])
		block.Time = time.Unix(int64(binary.LittleEndian.Uint32(blocksBytes[24:28])), 0)
		block.Pad3 = binary.LittleEndian.Uint32(blocksBytes[28:32])
		blockData = make([]byte, length-32)
		blockData.UnmarshalBytes(blocksBytes[32:length])
	} else {
		blockData = make([]byte, length-16)
		blockData.UnmarshalBytes(blocksBytes[16:length])
	}
	block.Data = &blockData
	return block, nil
}

func isValidHeader(header []byte) bool {
	preamble := binary.LittleEndian.Uint64(header[0:8])
	return (preamble == 0) || (preamble == 0xe2465dff41a05252)
}

func strictIsValidHeader(header []byte) bool {
	preamble := binary.LittleEndian.Uint64(header[0:8])
	return preamble == 0xe2465dff41a05252
}

// DiscardDataUntilValid will trim off invalid data on the buffered input
// until it reaches a header that is valid or the buffer has insufficient data.
// This is useful for when the input stream has been corrupted with some invalid bytes.
func (d *Decoder) DiscardDataUntilValid() {
	for {
		header, err := d.bufReader.Peek(28)
		if err != nil {
			return
		}

		if strictIsValidHeader(header) {
			return
		}
		_, _ = d.bufReader.Discard(1)
	}
}
