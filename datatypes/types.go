package datatypes

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"reflect"

	"github.com/ff14wed/xivnet"
)

var inTypeRegistry = make(map[uint16]xivnet.BlockData)
var outTypeRegistry = make(map[uint16]xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	AddStatusOpcode       = 0x141 // Unchanged for 4.5
	ActionOpcode          = 0x154 // Updated for 4.5
	AoEAction8Opcode      = 0x157 // Updated for 4.5
	AoEAction16Opcode     = 0x158 // Updated for 4.5
	AoEAction24Opcode     = 0x159 // Updated for 4.5
	AoEAction32Opcode     = 0x15A // Updated for 4.5
	CastingOpcode         = 0x17C // Updated for 4.5
	CraftStateOpcode      = 0x1AF // Updated for 4.5
	EquipChangeOpcode     = 0x18B // Updated for 4.5
	GaugeOpcode           = 0x29A // Updated for 4.5
	HateRankingOpcode     = 0x17F // Updated for 4.5
	HateListOpcode        = 0x180 // Updated for 4.5
	InitZoneOpcode        = 0x19A // Updated for 4.5
	MapChangeOpcode       = 0x299 // Updated for 4.5
	MarkerOpcode          = 0x267 // Updated for 4.5
	MountOpcode           = 0x1E8 // Updated for 4.5
	MovementOpcode        = 0x178 // Updated for 4.5
	NotifyOpcode          = 0x142 // Unchanged 4.5
	Notify3Opcode         = 0x143 // Unchanged 4.5
	Notify4Opcode         = 0x144 // Unchanged 4.5
	NPCSpawnOpcode        = 0x176 // Updated for 4.5
	NPCSpawn2Opcode       = 0x177 // Updated for 4.5
	PerformOpcode         = 0x2A5 // Updated for 4.5
	PlayerSpawnOpcode     = 0x175 // Updated for 4.5
	RemoveEntityOpcode    = 0x191 // Unchanged 4.5
	SetPosOpcode          = 0x17A // Updated for 4.5
	UpdateHPMPTPOpcode    = 0x145 // Unchanged 4.5
	UpdateStatusesOpcode  = 0x151 // Updated for 4.5
	WeatherChangeOpcode   = 0x205 // Updated for 4.5
	XWorldPartyListOpcode = 0xA1  // Updated 4.18
)

var _ = registerInBlockData(AddStatusOpcode, new(AddStatus))
var _ = registerInBlockData(ActionOpcode, new(Action))
var _ = registerInBlockData(AoEAction8Opcode, new(AoEAction8))
var _ = registerInBlockData(AoEAction16Opcode, new(AoEAction16))
var _ = registerInBlockData(AoEAction24Opcode, new(AoEAction24))
var _ = registerInBlockData(AoEAction32Opcode, new(AoEAction32))
var _ = registerInBlockData(CastingOpcode, new(Casting))
var _ = registerInBlockData(CraftStateOpcode, new(CraftState))
var _ = registerInBlockData(EquipChangeOpcode, new(EquipChange))
var _ = registerInBlockData(GaugeOpcode, new(Gauge))
var _ = registerInBlockData(HateListOpcode, new(HateList))
var _ = registerInBlockData(HateRankingOpcode, new(HateRanking))
var _ = registerInBlockData(InitZoneOpcode, new(InitZone))
var _ = registerInBlockData(MapChangeOpcode, new(MapChange))
var _ = registerInBlockData(MarkerOpcode, new(Marker))
var _ = registerInBlockData(MountOpcode, new(Mount))
var _ = registerInBlockData(MovementOpcode, new(Movement))
var _ = registerInBlockData(NotifyOpcode, new(Notify))
var _ = registerInBlockData(Notify3Opcode, new(Notify3))
var _ = registerInBlockData(Notify4Opcode, new(Notify4))
var _ = registerInBlockData(NPCSpawnOpcode, new(NPCSpawn))
var _ = registerInBlockData(NPCSpawn2Opcode, new(NPCSpawn2))
var _ = registerInBlockData(PerformOpcode, new(Perform))
var _ = registerInBlockData(PlayerSpawnOpcode, new(PlayerSpawn))
var _ = registerInBlockData(RemoveEntityOpcode, new(RemoveEntity))
var _ = registerInBlockData(SetPosOpcode, new(SetPos))
var _ = registerInBlockData(UpdateHPMPTPOpcode, new(UpdateHPMPTP))
var _ = registerInBlockData(UpdateStatusesOpcode, new(UpdateStatuses))
var _ = registerInBlockData(WeatherChangeOpcode, new(WeatherChange))

// var _ = registerInBlockData(XWorldPartyListOpcode, new(XWorldPartyList))

// Opcodes that define the datatypes of outgoing (to server) network blocks
const (
	MyActionOpcode      = 0x138 // Updated for 4.5
	MyMovementOpcode    = 0x13F // Updated for 4.5
	MyMovement2Opcode   = 0x17E // Updated for 4.5
	MyPerformOpcode     = 0x188 // Updated for 4.5
	BeginCraftingOpcode = 0x15D // Updated for 4.5
)

var _ = registerOutBlockData(MyActionOpcode, new(MyAction))
var _ = registerOutBlockData(MyMovementOpcode, new(MyMovement))
var _ = registerOutBlockData(MyMovement2Opcode, new(MyMovement2))
var _ = registerOutBlockData(MyPerformOpcode, new(Perform))
var _ = registerOutBlockData(BeginCraftingOpcode, new(BeginCrafting))

// NewBlockData is a factory for BlockData that uses the opcode to
// determine which BlockData to create
func NewBlockData(opcode uint16, isOut bool) xivnet.BlockData {
	r := inTypeRegistry
	if isOut {
		r = outTypeRegistry
	}
	template, ok := r[opcode]
	if !ok {
		return nil
	}
	rt := reflect.TypeOf(template)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	return reflect.New(rt).Interface().(xivnet.BlockData)
}

// UnmarshalBlockBytes converts raw bytes to this block data struct
func UnmarshalBlockBytes(data []byte, block xivnet.BlockData) error {
	if binary.Size(block) != len(data) {
		return fmt.Errorf("length mismatch: %d != %d", len(data), binary.Size(block))
	}
	return binary.Read(bytes.NewReader(data), binary.LittleEndian, block)
}

func registerInBlockData(opcode uint16, blockData xivnet.BlockData) struct{} {
	gob.Register(blockData)
	inTypeRegistry[opcode] = blockData
	return struct{}{}
}

func registerOutBlockData(opcode uint16, blockData xivnet.BlockData) struct{} {
	gob.Register(blockData)
	outTypeRegistry[opcode] = blockData
	return struct{}{}
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
	bd := NewBlockData(block.Header.Opcode, isOut)
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
