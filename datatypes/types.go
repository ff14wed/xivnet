package datatypes

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/ff14wed/xivnet/v3"
)

const UndefinedOpcode = 0xFFFF

func newBlockData(opcode uint16, registry map[uint16]func() xivnet.BlockData) xivnet.BlockData {
	factory, ok := registry[opcode]
	if !ok {
		return nil
	}
	return factory()
}

// NewBlockData is a factory for BlockData that uses the opcode to
// determine which BlockData to create
func NewBlockData(opcode uint16, isOut bool) xivnet.BlockData {
	r := inTypeRegistry
	if isOut {
		r = outTypeRegistry
	}
	return newBlockData(opcode, r)
}

type BlockUnmarshaler interface {
	UnmarshalBytes(data []byte) error
}

// UnmarshalBlockBytes converts raw bytes to this block data struct
func UnmarshalBlockBytes(data []byte, block xivnet.BlockData) error {
	if b, matches := block.(BlockUnmarshaler); matches {
		return b.UnmarshalBytes(data)
	}
	if binary.Size(block) != len(data) {
		return fmt.Errorf("length mismatch: %d != %d", len(data), binary.Size(block))
	}
	return binary.Read(bytes.NewReader(data), binary.LittleEndian, block)
}

// ParseBlock takes in raw, unparsed blocks and returns a parsed block if
// possible.
// isOut toggles whether we should parse this block as an outgoing block (sent
// packet) or as an incoming block (recv'd packet)
func ParseBlock(block *xivnet.Block, isOut bool) (*xivnet.Block, error) {
	data, ok := block.Data.(*xivnet.GenericBlockData)
	if !ok {
		return block, nil
	}
	blockBytes, _ := data.MarshalBytes()

	var bd xivnet.BlockData

	// This is kind of a hack to check if a packet is a chat type since it's
	// impossible to otherwise obtain this information without capturing the open
	// of the game connections.
	if block.ServerID == 0 && block.SubjectID != 0 {
		bd = NewChatBlockData(block.Opcode, isOut)
	} else {
		bd = NewBlockData(block.Opcode, isOut)
	}

	if bd == nil {
		return block, nil
	}
	err := UnmarshalBlockBytes(blockBytes, bd)
	if err != nil {
		return block, err
	}
	newB := *block
	newB.Data = bd
	return &newB, nil
}
