package datatypes

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/ff14wed/xivnet/v3"
)

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)
var outTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x10B // Updated for 5.15
	InitZoneOpcode      = 0x1BA // Updated for 5.15
	ControlOpcode       = 0x12F // Updated for 5.15
	ControlSelfOpcode   = 0x201 // Updated for 5.15
	ControlTargetOpcode = 0x1BE // Updated for 5.15
	RemoveEntityOpcode  = 0x32B // Updated for 5.15
	UpdateHPMPTPOpcode  = 0x75  // Updated for 5.15

	UpdateStatusesOpcode = 0x263 // Updated for 5.15
	// UpdateStatusesEurekaOpcode = 0x1C2 // Updated for 5.11
	UpdateStatusesBossOpcode = 0x312 // Updated for 5.15

	ActionOpcode      = 0x2AA // Updated for 5.15
	AoEAction8Opcode  = 0xB3  // Updated for 5.15
	AoEAction16Opcode = 0xE6  // Updated for 5.15
	AoEAction24Opcode = 0x10A // Updated for 5.15
	AoEAction32Opcode = 0x1C8 // Updated for 5.15

	PlayerSpawnOpcode = 0xDC  // Updated for 5.15
	NPCSpawnOpcode    = 0x219 // Updated for 5.15
	// NPCSpawn2Opcode   = 0x137 // Updated for 5.11

	MovementOpcode = 0x1AD // Updated for 5.15
	SetPosOpcode   = 0x296 // Updated for 5.15

	CastingOpcode = 0x1EC // Updated for 5.15

	HateRankingOpcode = 0x379 // Updated for 5.15
	HateListOpcode    = 0x351 // Updated for 5.15

	EquipChangeOpcode = 0x29B // Updated for 5.15

	// EventPlayOpcode         = 0x3B6 // Updated for 5.11
	// EventPlay2Opcode        = 0xA8  // Updated for 5.11
	// DirectorPlaySceneOpcode = 0x150 // Updated for 5.11

	MountOpcode = 0x3C0 // Updated for 5.15

	// WeatherChangeOpcode = 0x2FB // Updated for 5.11

	// WaymarkOpcode = 0x272 // Updated for 5.0

	PrepareZoningOpcode = 0x2B6 // Updated for 5.15

	GaugeOpcode = 0x337 // Updated for 5.15
	// PerformOpcode = 0x2A5 // Updated for 4.5

	// XWorldPartyListOpcode = 0xA1 // Updated 4.18
)

func init() {
	registerInBlockFactory(EffectResultOpcode, func() xivnet.BlockData { return new(EffectResult) })
	registerInBlockFactory(InitZoneOpcode, func() xivnet.BlockData { return new(InitZone) })
	registerInBlockFactory(ControlOpcode, func() xivnet.BlockData { return new(Control) })
	registerInBlockFactory(ControlSelfOpcode, func() xivnet.BlockData { return new(ControlSelf) })
	registerInBlockFactory(ControlTargetOpcode, func() xivnet.BlockData { return new(ControlTarget) })
	registerInBlockFactory(RemoveEntityOpcode, func() xivnet.BlockData { return new(RemoveEntity) })
	registerInBlockFactory(UpdateHPMPTPOpcode, func() xivnet.BlockData { return new(UpdateHPMPTP) })

	registerInBlockFactory(UpdateStatusesOpcode, func() xivnet.BlockData { return new(UpdateStatuses) })
	// registerInBlockFactory(UpdateStatusesEurekaOpcode, func() xivnet.BlockData { return new(UpdateStatusesEureka) })
	registerInBlockFactory(UpdateStatusesBossOpcode, func() xivnet.BlockData { return new(UpdateStatusesBoss) })

	registerInBlockFactory(ActionOpcode, func() xivnet.BlockData { return new(Action) })
	registerInBlockFactory(AoEAction8Opcode, func() xivnet.BlockData { return new(AoEAction8) })
	registerInBlockFactory(AoEAction16Opcode, func() xivnet.BlockData { return new(AoEAction16) })
	registerInBlockFactory(AoEAction24Opcode, func() xivnet.BlockData { return new(AoEAction24) })
	registerInBlockFactory(AoEAction32Opcode, func() xivnet.BlockData { return new(AoEAction32) })

	registerInBlockFactory(PlayerSpawnOpcode, func() xivnet.BlockData { return new(PlayerSpawn) })
	registerInBlockFactory(NPCSpawnOpcode, func() xivnet.BlockData { return new(NPCSpawn) })
	// registerInBlockFactory(NPCSpawn2Opcode, func() xivnet.BlockData { return new(NPCSpawn2) })

	registerInBlockFactory(MovementOpcode, func() xivnet.BlockData { return new(Movement) })
	registerInBlockFactory(SetPosOpcode, func() xivnet.BlockData { return new(SetPos) })
	registerInBlockFactory(CastingOpcode, func() xivnet.BlockData { return new(Casting) })

	registerInBlockFactory(HateRankingOpcode, func() xivnet.BlockData { return new(HateRanking) })
	registerInBlockFactory(HateListOpcode, func() xivnet.BlockData { return new(HateList) })

	registerInBlockFactory(EquipChangeOpcode, func() xivnet.BlockData { return new(EquipChange) })

	// registerInBlockFactory(EventPlayOpcode, func() xivnet.BlockData { return new(EventPlay) })
	// registerInBlockFactory(EventPlay2Opcode, func() xivnet.BlockData { return new(EventPlay2) })
	// registerInBlockFactory(DirectorPlaySceneOpcode, func() xivnet.BlockData { return new(DirectorPlayScene) })

	registerInBlockFactory(MountOpcode, func() xivnet.BlockData { return new(Mount) })

	// registerInBlockFactory(WeatherChangeOpcode, func() xivnet.BlockData { return new(WeatherChange) })

	// registerInBlockFactory(WaymarkOpcode, func() xivnet.BlockData { return new(Marker) })
	registerInBlockFactory(PrepareZoningOpcode, func() xivnet.BlockData { return new(PrepareZoning) })

	registerInBlockFactory(GaugeOpcode, func() xivnet.BlockData { return new(Gauge) })
	// registerInBlockFactory(PerformOpcode, func() xivnet.BlockData { return new(Perform) })

	// registerInBlockFactory(XWorldPartyListOpcode, new(XWorldPartyList))
}

// Opcodes that define the datatypes of outgoing (to server) network blocks
const (
	EgressClientTriggerOpcode = 0x371 // Updated for 5.15

	EgressMovementOpcode         = 0x355 // Updated for 5.15
	EgressInstanceMovementOpcode = 0x32B // Updated for 5.15

	// EgressPerformOpcode    = 0x18B // Updated for 5.0
	// EgressCraftEventOpcode = 0x299 // Updated for 5.11
)

func init() {
	registerOutBlockFactory(EgressClientTriggerOpcode, func() xivnet.BlockData { return new(EgressClientTrigger) })

	registerOutBlockFactory(EgressMovementOpcode, func() xivnet.BlockData { return new(EgressMovement) })
	registerOutBlockFactory(EgressInstanceMovementOpcode, func() xivnet.BlockData { return new(EgressInstanceMovement) })

	// registerOutBlockFactory(EgressPerformOpcode, func() xivnet.BlockData { return new(Perform) })

	// registerOutBlockFactory(EgressCraftEventOpcode, func() xivnet.BlockData { return new(EgressCraftEvent) })
}

// NewBlockData is a factory for BlockData that uses the opcode to
// determine which BlockData to create
func NewBlockData(opcode uint16, isOut bool) xivnet.BlockData {
	r := inTypeRegistry
	if isOut {
		r = outTypeRegistry
	}
	factory, ok := r[opcode]
	if !ok {
		return nil
	}
	return factory()
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

func registerInBlockFactory(opcode uint16, factory func() xivnet.BlockData) {
	inTypeRegistry[opcode] = factory
}

func registerOutBlockFactory(opcode uint16, factory func() xivnet.BlockData) {
	outTypeRegistry[opcode] = factory
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
	bd := NewBlockData(block.Opcode, isOut)
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
