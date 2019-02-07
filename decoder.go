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

// Decoder implements an FFXIV frame decoder
type Decoder struct {
	buf     []byte
	bufSize int
}

// NewDecoder creates a new instance of a decoder
func NewDecoder(bufSize int) *Decoder {
	return &Decoder{
		buf:     make([]byte, bufSize),
		bufSize: bufSize,
	}
}

// CheckHeader checks to see whether or not the data in the buffer has a
// valid header
func (d *Decoder) CheckHeader(buf *bufio.Reader) ([]byte, error) {
	// Validation of lengths and stuff
	header, err := buf.Peek(28)
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

// Decode returns a single decoded FFXIV frame from the packet data stored in
// buf
func (d *Decoder) Decode(buf *bufio.Reader) (*Frame, error) {
	header, err := d.CheckHeader(buf)
	if err != nil {
		return nil, err
	}
	length := binary.LittleEndian.Uint32(header[24:])
	if length > uint32(d.bufSize) {
		return nil, InvalidFrameLengthError{length: length, maxLength: d.bufSize}
	}
	intLength := int(length)
	_, err = buf.Peek(intLength)
	if err != nil {
		return nil, EOFError{
			operation:       "peeking data",
			attemptedLength: intLength,
			wrapped:         err,
		}
	}
	n, err := buf.Read(d.buf[:length])
	if err != nil {
		return nil, EOFError{
			operation:       "reading data",
			attemptedLength: intLength,
			wrapped:         err,
		}
	}
	if n != int(length) {
		return nil, MismatchedReadLengthsError{
			readLength: n, expectedLength: intLength,
		}
	}
	f, err := decodeFrame(d.buf[:length], d.buf[length:], length)
	if err != nil {
		return nil, DecodingError{
			wrapped:   err,
			debugData: hex.EncodeToString(d.buf[:length]),
		}
	}
	return f, nil
}

func decodeFrame(frameBytes []byte, blockBuffer []byte, length uint32) (*Frame, error) {
	// Build the frame
	frame := &Frame{}
	copy(frame.Header[:], frameBytes[0:16])
	msecSinceEpoch := time.Duration(binary.LittleEndian.Uint64(frameBytes[16:24])) * time.Millisecond
	frame.Time = time.Unix(0, 0).Add(msecSinceEpoch)
	frame.Length = length
	frame.Reserved1 = binary.LittleEndian.Uint16(frameBytes[28:30])
	frame.NumBlocks = binary.LittleEndian.Uint16(frameBytes[30:32])
	frame.Compression = binary.LittleEndian.Uint16(frameBytes[32:34])
	frame.Reserved2 = binary.LittleEndian.Uint32(frameBytes[34:38])
	frame.Reserved3 = binary.LittleEndian.Uint16(frameBytes[38:40])

	blockData := frameBytes[40:length]
	if frame.Compression != 0 && frame.Compression != 1 {
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
	var blockData GenericBlockData
	if length < 32 {
		blockData = make([]byte, length-4)
		_ = blockData.UnmarshalBytes(blocksBytes[4:length])
	} else {
		block.Header.SubjectID = binary.LittleEndian.Uint32(blocksBytes[4:8])
		block.Header.CurrentID = binary.LittleEndian.Uint32(blocksBytes[8:12])
		block.Header.U1 = binary.LittleEndian.Uint32(blocksBytes[12:16])
		block.Header.U2 = binary.LittleEndian.Uint16(blocksBytes[16:18])
		block.Header.Opcode = binary.LittleEndian.Uint16(blocksBytes[18:20])
		block.Header.Route = binary.LittleEndian.Uint32(blocksBytes[20:24])
		block.Header.Time = time.Unix(int64(binary.LittleEndian.Uint32(blocksBytes[24:28])), 0)
		block.Header.U4 = binary.LittleEndian.Uint32(blocksBytes[28:32])
		blockData = make([]byte, length-32)
		_ = blockData.UnmarshalBytes(blocksBytes[32:length])
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
func (d *Decoder) DiscardDataUntilValid(buf *bufio.Reader) {
	for {
		header, err := buf.Peek(28)
		if err != nil {
			return
		}

		if strictIsValidHeader(header) {
			return
		}
		_, _ = buf.Discard(1)
	}
}
