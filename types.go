package xivnet

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"strings"
	"sync"
	"time"
)

// These types define the FFXIV packet structures. Some of the fields (like
// length or compression) aren't really necessary, but they codify some of the
// fields that need to be parsed in a packet.

var mbufPool *marshalBufferPool

type marshalBufferPool struct {
	pool *sync.Pool
}

func newPool() *marshalBufferPool {
	return &marshalBufferPool{
		pool: new(sync.Pool),
	}
}

func (m *marshalBufferPool) Get() *bytes.Buffer {
	b := m.pool.Get()
	if b == nil {
		return new(bytes.Buffer)
	}
	return b.(*bytes.Buffer)
}

func (m *marshalBufferPool) Put(b *bytes.Buffer) {
	b.Reset()
	m.pool.Put(b)
}

// Header defines the type for a 16 byte array
type Header [16]byte

// MarshalJSON returns the marshaled version of the frame header
func (h *Header) MarshalJSON() ([]byte, error) {
	return bytesToByteString(h[:])
}

// UnmarshalJSON returns the unmarshaled version of the frame header
func (h *Header) UnmarshalJSON(data []byte) error {
	newB, err := byteStringToBytes(data)
	if err != nil {
		return err
	}
	copy(h[:], newB)
	return nil
}

// BlockData defines the interface for the XIVWS representation of FFXIV
// block data
type BlockData interface{}

// MarshalBlockBytes returns the byte representation of the block data
func MarshalBlockBytes(block BlockData) ([]byte, error) {
	byteBuf := new(bytes.Buffer)
	err := binary.Write(byteBuf, binary.LittleEndian, block)
	if err != nil {
		return nil, err
	}
	return byteBuf.Bytes(), nil
}

// GenericBlockData defines the type for a variable length byte slice
type GenericBlockData []byte

var _ BlockData = new(GenericBlockData)

// GenericBlockDataFromBytes is a helper that creates a GenericBlockData from
// raw bytes
func GenericBlockDataFromBytes(blockData []byte) *GenericBlockData {
	var g GenericBlockData = blockData
	return &g
}

// MarshalJSON returns the marshaled version of the BlockHeader
func (b *GenericBlockData) MarshalJSON() ([]byte, error) {
	return bytesToByteString(*b)
}

// UnmarshalJSON returns the unmarshaled version of the BlockHeader
func (b *GenericBlockData) UnmarshalJSON(data []byte) error {
	newB, err := byteStringToBytes(data)
	if err != nil {
		return err
	}
	*b = newB
	return nil
}

// MarshalBytes returns the byte representation of the block data
func (b *GenericBlockData) MarshalBytes() ([]byte, error) {
	return *b, nil
}

// UnmarshalBytes sets the block data to the provided raw bytes
func (b *GenericBlockData) UnmarshalBytes(data []byte) error {
	copy(*b, data)
	return nil
}

// Length returns the length of the block data
func (b *GenericBlockData) Length() uint32 {
	return uint32(len(*b))
}

func bytesToByteString(b []byte) ([]byte, error) {
	n := len(b)
	if n == 0 {
		return []byte(`""`), nil
	}
	if mbufPool == nil {
		mbufPool = newPool()
	}
	marshalBuffer := mbufPool.Get()
	defer mbufPool.Put(marshalBuffer)

	marshalBuffer.WriteRune('"')
	if b[0] < 0x10 {
		marshalBuffer.WriteRune('0')
	}
	marshalBuffer.WriteString(strconv.FormatUint(uint64(b[0]), 16))
	for i := 1; i < n; i++ {
		marshalBuffer.WriteRune(' ')
		if b[i] < 0x10 {
			marshalBuffer.WriteRune('0')
		}
		marshalBuffer.WriteString(strconv.FormatUint(uint64(b[i]), 16))
	}
	marshalBuffer.WriteRune('"')
	return marshalBuffer.Bytes(), nil
}

func byteStringToBytes(b []byte) ([]byte, error) {
	n := len(b)
	encodedHex := strings.Replace(string(b[1:n-1]), " ", "", -1)
	return hex.DecodeString(encodedHex)
}

// Frame defines the structure for an FFXIV Frame
type Frame struct {
	Header      Header
	Time        time.Time
	Length      uint32
	Reserved1   uint16
	NumBlocks   uint16
	Compression uint16
	Reserved2   uint32
	Reserved3   uint16
	Blocks      []*Block

	compressedBlockData []byte
}

// Block defines the structure of a block in an FFXIV frame
type Block struct {
	Length uint32
	Header BlockHeader
	Data   BlockData
}

// BlockHeader defines the type for a 28 byte array
type BlockHeader struct {
	SubjectID uint32
	CurrentID uint32
	U1        uint32
	U2        uint16
	Opcode    uint16
	Route     uint32
	Time      time.Time
	U4        uint32
}
