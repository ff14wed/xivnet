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
	AddStatusOpcode       = 0x141 // Unchanged for 4.3
	ActionOpcode          = 0x151 // Updated for 4.3
	AoEAction8Opcode      = 0x154 // Updated for 4.3
	AoEAction16Opcode     = 0x155 // Updated for 4.3
	AoEAction24Opcode     = 0x156 // Updated for 4.3
	AoEAction32Opcode     = 0x157 // Updated for 4.3
	CastingOpcode         = 0x178 // Updated for 4.3
	EquipChangeOpcode     = 0x186 // Updated for 4.3
	GaugeOpcode           = 0x292 // Updated for 4.3
	HateListOpcode        = 0x17C // Updated for 4.3
	HateRankingOpcode     = 0x17B // Updated for 4.3
	InitZoneOpcode        = 0x19A // Unchanged 4.3
	MapChangeOpcode       = 0x291 // Updated for 4.3
	MountOpcode           = 0x1E3 // Updated for 4.3
	MovementOpcode        = 0x174 // Updated for 4.3
	NotifyOpcode          = 0x142 // Unchanged 4.3
	Notify3Opcode         = 0x143 // Unchanged 4.3
	Notify4Opcode         = 0x144 // Unchanged 4.3
	NPCSpawnOpcode        = 0x173 // Updated for 4.3
	PerformOpcode         = 0x286 // Updated for 4.3
	PlayerSpawnOpcode     = 0x172 // Updated for 4.3
	RemoveEntityOpcode    = 0x191 // Unchanged 4.3
	SetPosOpcode          = 0x176 // Updated for 4.3
	UpdateHPMPTPOpcode    = 0x145 // Unchanged 4.3
	UpdateStatusesOpcode  = 0x14E // Updated for 4.3
	WeatherChangeOpcode   = 0x200 // Updated for 4.3
	XWorldPartyListOpcode = 0xA1  // Updated 4.18
)

var _ = registerInBlockData(AddStatusOpcode, new(AddStatus))

// var _ = registerInBlockData(ActionOpcode, new(Action))
// var _ = registerInBlockData(AoEAction8Opcode, new(AoEAction8))
// var _ = registerInBlockData(AoEAction16Opcode, new(AoEAction16))
// var _ = registerInBlockData(AoEAction24Opcode, new(AoEAction24))
// var _ = registerInBlockData(AoEAction32Opcode, new(AoEAction32))
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
var _ = registerInBlockData(Notify4Opcode, new(Notify4))
var _ = registerInBlockData(NPCSpawnOpcode, new(NPCSpawn))
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
	MyActionOpcode    = 0x13D // Updated for 4.3
	MyMovementOpcode  = 0x144 // Updated for 4.3
	MyMovement2Opcode = 0x183 // Updated for 4.3
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
