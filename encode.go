package xivnet

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"io"
	"time"
)

// CorrectTimestamps corrects the timestamps in the FFXIV frame header and its
// blocks
func (f *Frame) CorrectTimestamps(timestamp time.Time) {
	f.Time = timestamp
	for _, b := range f.Blocks {
		b.Time = timestamp
	}
}

// correctLength returns the correct length of the frame header plus all its
// blocks. This method assumes the blocks are uncompressed.
func (f Frame) correctLength() uint32 {
	var blocksLength uint32
	for _, b := range f.Blocks {
		blocksLength += b.CorrectLength()
	}
	return 40 + blocksLength
}

// CorrectLength returns the true length of the block
func (b Block) CorrectLength() uint32 {
	if b.Type == BlockTypeIPC {
		return uint32(32 + binary.Size(b.Data))
	}
	return uint32(16 + binary.Size(b.Data))
}

// Encode writes the byte representation of the block to the output writer
func (b Block) Encode(w io.Writer) error {
	correctLength := b.CorrectLength()
	if (correctLength < 32 && b.Type == 3) || correctLength < 16 {
		// Error should never happen
		return errors.New("Block length is too small")
	}
	headerLength := 16
	buf := make([]byte, correctLength)
	binary.LittleEndian.PutUint32(buf[0:4], correctLength)
	binary.LittleEndian.PutUint32(buf[4:8], b.SubjectID)
	binary.LittleEndian.PutUint32(buf[8:12], b.CurrentID)
	binary.LittleEndian.PutUint16(buf[12:14], b.Type)
	binary.LittleEndian.PutUint16(buf[14:16], b.Pad1)
	if b.Type == 3 {
		headerLength = 32
		binary.LittleEndian.PutUint16(buf[16:18], b.Reserved)
		binary.LittleEndian.PutUint16(buf[18:20], b.Opcode)
		binary.LittleEndian.PutUint16(buf[20:22], b.Pad2)
		binary.LittleEndian.PutUint16(buf[22:24], b.ServerID)
		time := uint32(b.Time.Unix())
		binary.LittleEndian.PutUint32(buf[24:28], time)
		binary.LittleEndian.PutUint32(buf[28:32], b.Pad3)
	}

	var blockData []byte
	var err error
	switch v := b.Data.(type) {
	case *GenericBlockData:
		blockData, err = v.MarshalBytes()
	default:
		blockData, err = MarshalBlockBytes(v)
	}
	if err != nil {
		return err
	}
	copy(buf[headerLength:correctLength], blockData)

	_, err = w.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

// compressBlocks returns all of the blocks compressed in byte format
func (f Frame) compressBlocks() ([]byte, error) {
	buf := new(bytes.Buffer)

	tmpBuf := new(bytes.Buffer)
	for _, b := range f.Blocks {
		err := b.Encode(tmpBuf)
		if err != nil {
			return nil, err
		}
	}

	zlibWriter := zlib.NewWriter(buf)
	_, err := tmpBuf.WriteTo(zlibWriter)
	if err != nil {
		return nil, err
	}
	zlibWriter.Close()

	return buf.Bytes(), nil
}

// Encode writes the byte representation of the frame to the output writer
// with the given timestamp
func (f Frame) Encode(w io.Writer, timestamp time.Time, compress bool) error {
	f.CorrectTimestamps(timestamp)
	f.Length = f.correctLength()
	f.Count = uint16(len(f.Blocks))
	f.Reserved1 = 1

	var compressedBlockData []byte
	if compress {
		var err error
		compressedBlockData, err = f.compressBlocks()
		if err != nil {
			return err
		}
		f.Compression = 1
		f.Length = uint32(40 + len(compressedBlockData))
	} else {
		f.Compression = 0
	}

	buf := make([]byte, 40)
	copy(buf[0:16], f.Preamble[:])
	time := uint64(f.Time.UnixNano() / 1000000)
	binary.LittleEndian.PutUint64(buf[16:24], time)
	binary.LittleEndian.PutUint32(buf[24:28], f.Length)
	binary.LittleEndian.PutUint16(buf[28:30], f.ConnectionType)
	binary.LittleEndian.PutUint16(buf[30:32], f.Count)
	buf[32] = f.Reserved1
	buf[33] = f.Compression
	binary.LittleEndian.PutUint32(buf[34:38], f.Reserved2)
	binary.LittleEndian.PutUint16(buf[38:40], f.Reserved3)

	_, err := w.Write(buf)
	if err != nil {
		return err
	}

	if compress {
		_, err := w.Write(compressedBlockData)
		if err != nil {
			return err
		}
	} else {
		for _, b := range f.Blocks {
			err := b.Encode(w)
			if err != nil {
				return err
			}
		}

	}
	return nil
}
