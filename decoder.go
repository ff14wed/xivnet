package xivnet

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"time"
)

// Decoder implements an FFXIV frame decoder. It is not thread-safe, so consumers
// need to be sure not to use it concurrently.
type Decoder struct {
	frameReader *bufio.Reader
	blockReader *bufio.Reader
}

// NewDecoder creates a new instance of a frame decoder.
// bufSize controls the total size of the buffer used to store a single
// frame. It's recommended to keep this value at least 8192, since this value
// generally works for all frames (as far as I can tell).
func NewDecoder(r io.Reader, bufSize int) *Decoder {
	return &Decoder{
		frameReader: bufio.NewReaderSize(r, bufSize),
		blockReader: bufio.NewReaderSize(r, 4096),
	}
}

// CheckHeader checks to see whether or not the data in the buffer has a
// valid header
func (d *Decoder) CheckHeader() ([]byte, error) {
	if 28 > d.frameReader.Size() {
		return nil, InvalidFrameLengthError{length: 28, maxLength: d.frameReader.Size()}
	}
	// Validation that the frame at least has
	header, err := d.frameReader.Peek(28)
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
// If there is some sort of decoding error, NextFrame will call
// DiscardDataUntilValid in order to recover from corrupt data in the byte
// stream.
func (d *Decoder) NextFrame() (*Frame, error) {
	header, err := d.CheckHeader()
	if err != nil {
		return nil, err
	}
	length := binary.LittleEndian.Uint32(header[24:])
	if length > uint32(d.frameReader.Size()) {
		return nil, InvalidFrameLengthError{length: length, maxLength: d.frameReader.Size()}
	}
	intLength := int(length)
	frameBytes, err := d.frameReader.Peek(intLength)
	if err != nil {
		return nil, EOFError{
			operation:       "peeking data",
			attemptedLength: intLength,
			wrapped:         err,
		}
	}
	f, err := decodeFrame(frameBytes, d.blockReader, length)
	if err != nil {
		// If there is some sort of decoding error, let's
		// start discarding data until it's valid
		debugData := hex.EncodeToString(frameBytes)
		_, _ = d.frameReader.Discard(1)
		d.DiscardDataUntilValid()
		return nil, DecodingError{
			wrapped:   err,
			debugData: debugData,
		}
	}
	// We "read" intLength amount of data
	_, _ = d.frameReader.Discard(intLength)
	return f, nil
}

func decodeFrame(frameBytes []byte, blockReader *bufio.Reader, length uint32) (*Frame, error) {
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
		blockReader.Reset(r)
	} else {
		blockReader.Reset(bytes.NewReader(blockData))
	}

	if len(blockData) == 0 {
		return frame, nil
	}

	for {
		block, err := decodeBlock(blockReader)
		if err != nil {
			return nil, fmt.Errorf("error decoding blocks: %s", err.Error())
		}
		if block == nil {
			break
		}
		frame.Blocks = append(frame.Blocks, block)
	}
	return frame, nil
}

func decodeBlock(reader *bufio.Reader) (*Block, error) {
	lengthBytes, err := reader.Peek(4)
	if err == io.EOF {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	length := binary.LittleEndian.Uint32(lengthBytes)
	blockBytes, err := reader.Peek(int(length))
	actualLen := len(blockBytes)
	defer func() {
		// No matter what, we "read" actualLen amount of data
		_, _ = reader.Discard(actualLen)
	}()
	if err != nil {
		return nil, fmt.Errorf("not enough data: expected %d bytes, got %d", length, actualLen)
	}
	if actualLen < 16 {
		return nil, fmt.Errorf("not enough data: expected at least 16 bytes, got %d", actualLen)
	}
	block := &Block{}
	block.Length = length
	block.SubjectID = binary.LittleEndian.Uint32(blockBytes[4:8])
	block.CurrentID = binary.LittleEndian.Uint32(blockBytes[8:12])
	block.Type = binary.LittleEndian.Uint16(blockBytes[12:14])
	block.Pad1 = binary.LittleEndian.Uint16(blockBytes[14:16])
	var blockData GenericBlockData
	if block.Type == BlockTypeIPC {
		if actualLen < 32 {
			return nil, fmt.Errorf("not enough data: expected at least 32 bytes, got %d", actualLen)
		}
		block.Reserved = binary.LittleEndian.Uint16(blockBytes[16:18])
		block.Opcode = binary.LittleEndian.Uint16(blockBytes[18:20])
		block.Pad2 = binary.LittleEndian.Uint16(blockBytes[20:22])
		block.ServerID = binary.LittleEndian.Uint16(blockBytes[22:24])
		block.Time = time.Unix(int64(binary.LittleEndian.Uint32(blockBytes[24:28])), 0)
		block.Pad3 = binary.LittleEndian.Uint32(blockBytes[28:32])
		blockData = make([]byte, length-32)
		blockData.UnmarshalBytes(blockBytes[32:length])
	} else {
		blockData = make([]byte, length-16)
		blockData.UnmarshalBytes(blockBytes[16:length])
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
		header, err := d.frameReader.Peek(28)
		if err != nil {
			return
		}

		if strictIsValidHeader(header) {
			return
		}
		_, _ = d.frameReader.Discard(1)
	}
}
