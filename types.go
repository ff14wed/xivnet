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

// Preamble defines the type for a 16 byte array
type Preamble [16]byte

// MarshalJSON returns the marshaled version of the frame magic
func (p *Preamble) MarshalJSON() ([]byte, error) {
	return bytesToByteString(p[:])
}

// UnmarshalJSON returns the unmarshaled version of the frame magic
func (p *Preamble) UnmarshalJSON(data []byte) error {
	newB, err := byteStringToBytes(data)
	if err != nil {
		return err
	}
	copy(p[:], newB)
	return nil
}

// BlockData defines the interface for the XIVWS representation of FFXIV
// block data
type BlockData interface {
	IsBlockData()
}

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

func (GenericBlockData) IsBlockData() {}

// GenericBlockDataFromBytes is a helper that creates a GenericBlockData from
// raw bytes
func GenericBlockDataFromBytes(blockData []byte) *GenericBlockData {
	var g GenericBlockData = blockData
	return &g
}

// MarshalJSON returns the marshaled version of the generic block data
func (b *GenericBlockData) MarshalJSON() ([]byte, error) {
	return bytesToByteString(*b)
}

// UnmarshalJSON returns the unmarshaled version of the generic block data
func (b *GenericBlockData) UnmarshalJSON(data []byte) error {
	newB, err := byteStringToBytes(data)
	if err != nil {
		return err
	}
	*b = newB
	return nil
}

// MarshalBytes returns the byte representation of the generic block data
func (b *GenericBlockData) MarshalBytes() ([]byte, error) {
	return *b, nil
}

// UnmarshalBytes sets the block data to the provided raw bytes
func (b *GenericBlockData) UnmarshalBytes(data []byte) {
	copy(*b, data)
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

// Frame defines the structure for an FFXIV Frame.
// More details in Sapphire's `Network/CommonNetwork.h`.
// The Frame header is encoded as 40 bytes on the wire.
type Frame struct {
	Preamble       Preamble  // [0:16] - Used to identify the start of a frame
	Time           time.Time // [16:24] - Number of milliseconds since the Unix epoch
	Length         uint32    // [24:28] - Total frame size, including the header
	ConnectionType uint16    // [28:30] - Connection type (0 lobby, 1 zone, 2 chat)
	Count          uint16    // [30:32] - Number of blocks in this frame
	Reserved1      byte      // [32]    - Usually 1
	Compression    byte      // [33]    - 1 if compressed, 0 if not
	Reserved2      uint32    // [34:38]
	Reserved3      uint16    // [38:40]
	Blocks         []*Block
}

// These constants indicate the type of block.
const (
	BlockTypeSessionInit = 1
	BlockTypeSessionRecv = 2
	BlockTypeIPC         = 3
	BlockTypePing        = 7
	BlockTypePong        = 8
	BlockTypeEncryptInit = 9
	BlockTypeEncryptRecv = 10
)

// Block defines the structure of a block in an FFXIV frame.
// More details in Sapphire's `Network/CommonNetwork.h`.
type Block struct {
	Length    uint32 // [0:4] - Total block size, including the header
	SubjectID uint32 // [4:8] - The session ID that this block describes
	CurrentID uint32 // [8:12] - The session ID of the sender/receiver of this block
	Type      uint16 // [12:14] - The segment type
	Pad1      uint16 // [14:16]
	IPCHeader        // [16:32] if Type == BlockTypeIPC
	Data      BlockData
}

// IPCHeader defines the type for the IPC header of an FFXIV block.
// In cases other than SEGMENTTYPE_IPC, this header will not be present.
type IPCHeader struct {
	Reserved uint16    // [16:18] - 0x14
	Opcode   uint16    // [18:20] - Type of IPC message
	Pad2     uint16    // [20:22]
	ServerID uint16    // [22:24] - Server ID handling this message
	Time     time.Time // [24:28] - Number of seconds since Unix epoch
	Pad3     uint32    // [28:32]
}
