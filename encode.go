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
		b.Header.Time = timestamp
	}
}

// CorrectLength corrects the length in the FFXIV frame header. This method
// assumes all the block header lengths are correct.
func (f *Frame) CorrectLength() {
	var blocksLength uint32
	for _, b := range f.Blocks {
		b.CorrectLength()
		blocksLength += b.Length
	}
	f.Length = 40 + blocksLength
}

// CorrectLength computes the true length of the block and sets its Length field
func (b *Block) CorrectLength() {
	b.Length = uint32(32 + binary.Size(b.Data))
}

// Encode writes the byte representation of the block to the output writer
func (b *Block) Encode(w io.Writer) error {
	if b.Length < 32 {
		return errors.New("Block length is too small")
	}
	buf := make([]byte, b.Length)
	binary.LittleEndian.PutUint32(buf[0:4], b.Length)
	binary.LittleEndian.PutUint32(buf[4:8], b.Header.SubjectID)
	binary.LittleEndian.PutUint32(buf[8:12], b.Header.CurrentID)
	binary.LittleEndian.PutUint32(buf[12:16], b.Header.U1)
	binary.LittleEndian.PutUint16(buf[16:18], b.Header.U2)
	binary.LittleEndian.PutUint16(buf[18:20], b.Header.Opcode)
	binary.LittleEndian.PutUint32(buf[20:24], b.Header.U3)
	time := uint32(b.Header.Time.Unix())
	binary.LittleEndian.PutUint32(buf[24:28], time)
	binary.LittleEndian.PutUint32(buf[28:32], b.Header.U4)

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
	copy(buf[32:b.Length], blockData)

	_, err = w.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

// CompressBlocks prepares the frame for writing by saving an internal
// representation of the compressed block bytes
func (f *Frame) CompressBlocks() error {
	buf := new(bytes.Buffer)

	tmpBuf := new(bytes.Buffer)
	for _, b := range f.Blocks {
		err := b.Encode(tmpBuf)
		if err != nil {
			return err
		}
	}

	zlibWriter := zlib.NewWriter(buf)
	_, err := tmpBuf.WriteTo(zlibWriter)
	if err != nil {
		return err
	}
	zlibWriter.Close()

	f.compressedBlockData = buf.Bytes()
	f.Length = uint32(40 + len(f.compressedBlockData))
	return nil
}

// Encode writes the byte representation of the frame to the output writer
// with the given timestamp
func (f *Frame) Encode(w io.Writer, timestamp time.Time, compress bool) error {
	f.CorrectLength()
	f.CorrectTimestamps(timestamp)
	f.NumBlocks = uint16(len(f.Blocks))
	f.Compression = 0x0001

	if compress {
		f.CompressBlocks()
		f.Compression = 0x0101
	}

	buf := make([]byte, 40)
	copy(buf[0:16], f.Header[:])
	time := uint64(f.Time.UnixNano() / 1000000)
	binary.LittleEndian.PutUint64(buf[16:24], time)
	binary.LittleEndian.PutUint32(buf[24:28], f.Length)
	binary.LittleEndian.PutUint16(buf[28:30], f.Reserved1)
	binary.LittleEndian.PutUint16(buf[30:32], f.NumBlocks)
	binary.LittleEndian.PutUint16(buf[32:34], f.Compression)
	binary.LittleEndian.PutUint32(buf[34:38], f.Reserved2)
	binary.LittleEndian.PutUint16(buf[38:40], f.Reserved3)

	_, err := w.Write(buf)
	if err != nil {
		return err
	}

	if compress {
		_, err := w.Write(f.compressedBlockData)
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
