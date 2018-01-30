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
	AddStatusOpcode       = 0x141 // Unchanged for 4.2
	ActionOpcode          = 0x128 // Updated for 4.2
	AoEAction8Opcode      = 0x12B // Updated for 4.2
	AoEAction16Opcode     = 0x138 // Updated for 4.2
	AoEAction24Opcode     = 0x139 // Updated for 4.2
	AoEAction32Opcode     = 0x13A // Updated for 4.2
	CastingOpcode         = 0x162 // Updated for 4.2
	EquipChangeOpcode     = 0x170 // Updated for 4.2
	GaugeOpcode           = 0x27D // Updated for 4.2
	HateListOpcode        = 0x166 // Updated for 4.2
	HateRankingOpcode     = 0x165 // Updated for 4.2
	InitZoneOpcode        = 0x19A // Unchanged 4.2
	MapChangeOpcode       = 0x27C // Updated for 4.2
	MountOpcode           = 0x1CD // Updated for 4.2
	MovementOpcode        = 0x15E // Updated for 4.2
	NotifyOpcode          = 0x142 // Unchanged 4.2
	Notify3Opcode         = 0x143 // Unchanged 4.2
	NPCSpawnOpcode        = 0x15D // Updated for 4.2
	PerformOpcode         = 0x286 // Updated for 4.2
	PlayerSpawnOpcode     = 0x15C // Updated for 4.2
	RemoveEntityOpcode    = 0x191 // Unchanged 4.2
	SetPosOpcode          = 0x160 // Updated for 4.2
	TargetOpcode          = 0x144 // Unchanged 4.2
	UpdateHPMPTPOpcode    = 0x145 // Unchanged 4.2
	UpdateStatusesOpcode  = 0x125 // Updated 4.2
	WeatherChangeOpcode   = 0x1DD // DOESN'T WORK
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
var _ = registerInBlockData(MountOpcode, new(Mount))
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

// var _ = registerInBlockData(WeatherChangeOpcode, new(WeatherChange))

// var _ = registerInBlockData(XWorldPartyListOpcode, new(XWorldPartyList))

// Opcodes that define the datatypes of outgoing (to server) network blocks
const (
	MyActionOpcode    = 0x131 // Updated for 4.18
	MyMovementOpcode  = 0x138 // Updated for 4.2
	MyMovement2Opcode = 0x177 // Updated for 4.2
	MyPerformOpcode   = 0x180 // Updated for 4.18
)

var _ = registerOutBlockData(MyActionOpcode, new(MyAction))
var _ = registerOutBlockData(MyMovementOpcode, new(MyMovement))
var _ = registerOutBlockData(MyMovement2Opcode, new(MyMovement2))
var _ = registerOutBlockData(MyPerformOpcode, new(Perform))

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
