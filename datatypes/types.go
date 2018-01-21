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
	AddStatusOpcode       = 0x141 // Updated for 4.1
	ActionOpcode          = 0xFB  // Updated for 4.1
	AoEAction8Opcode      = 0xFE  // Updated for 4.1
	AoEAction16Opcode     = 0xFF  // Updated for 4.1
	AoEAction24Opcode     = 0x100 // Updated for 4.1
	AoEAction32Opcode     = 0x101 // Updated for 4.1
	CastingOpcode         = 0x123 // Updated for 4.1
	EquipChangeOpcode     = 0x13C // Updated for 4.1
	GaugeOpcode           = 0x249 // Updated for 4.1
	HateListOpcode        = 0x127 // Updated for 4.1
	HateRankingOpcode     = 0x126 // Updated for 4.1
	InitZoneOpcode        = 0x19A // Unchanged 4.1
	MapChangeOpcode       = 0x248 // Updated for 4.1
	MovementOpcode        = 0x11E // Updated for 4.1
	NotifyOpcode          = 0x142 // Unchanged 4.1
	Notify3Opcode         = 0x143 // Unchanged 4.1
	NPCSpawnOpcode        = 0x11D // Updated for 4.1
	PerformOpcode         = 0x252 // Updated for 4.15
	PlayerSpawnOpcode     = 0x11C // Updated for 4.1
	RemoveEntityOpcode    = 0x191 // Unchanged 4.1
	SetPosOpcode          = 0x120 // Updated for 4.1
	TargetOpcode          = 0x144 // Unchanged 4.1
	UpdateHPMPTPOpcode    = 0x145 // Unchanged 4.1
	UpdateStatusesOpcode  = 0xFA  // Updated 4.1
	XWorldPartyListOpcode = 0xA1  // Updated 4.18
)

var _ = registerInBlockData(AddStatusOpcode, new(AddStatus))
var _ = registerInBlockData(ActionOpcode, new(Action))
var _ = registerInBlockData(AoEAction8Opcode, new(AoEAction8))
var _ = registerInBlockData(AoEAction16Opcode, new(AoEAction16))
var _ = registerInBlockData(AoEAction24Opcode, new(AoEAction24))
var _ = registerInBlockData(AoEAction32Opcode, new(AoEAction32))
var _ = registerInBlockData(CastingOpcode, new(Casting))
var _ = registerInBlockData(EquipChangeOpcode, new(EquipChange))
var _ = registerInBlockData(GaugeOpcode, new(Gauge))
var _ = registerInBlockData(HateListOpcode, new(HateList))
var _ = registerInBlockData(HateRankingOpcode, new(HateRanking))
var _ = registerInBlockData(InitZoneOpcode, new(InitZone))
var _ = registerInBlockData(MapChangeOpcode, new(MapChange))
var _ = registerInBlockData(MovementOpcode, new(Movement))
var _ = registerInBlockData(NotifyOpcode, new(Notify))
var _ = registerInBlockData(Notify3Opcode, new(Notify3))
var _ = registerInBlockData(NPCSpawnOpcode, new(NPCSpawn))
var _ = registerInBlockData(PerformOpcode, new(Perform))
var _ = registerInBlockData(PlayerSpawnOpcode, new(PlayerSpawn))
var _ = registerInBlockData(RemoveEntityOpcode, new(RemoveEntity))
var _ = registerInBlockData(SetPosOpcode, new(SetPos))
var _ = registerInBlockData(TargetOpcode, new(Target))
var _ = registerInBlockData(UpdateHPMPTPOpcode, new(UpdateHPMPTP))
var _ = registerInBlockData(UpdateStatusesOpcode, new(UpdateStatuses))
var _ = registerInBlockData(XWorldPartyListOpcode, new(XWorldPartyList))

// Opcodes that define the datatypes of outgoing (to server) network blocks
const (
	MyActionOpcode    = 0x111 // Updated for 4.18
	MyMovementOpcode  = 0x118 // Updated for 4.1
	MyMovement2Opcode = 0x157 // Updated for 4.18
)

var _ = registerOutBlockData(MyActionOpcode, new(MyAction))
var _ = registerOutBlockData(MyMovementOpcode, new(MyMovement))
var _ = registerOutBlockData(MyMovement2Opcode, new(MyMovement2))

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
