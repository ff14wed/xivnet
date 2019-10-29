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
	// Opcodes that change rarely
	EffectResultOpcode  = 0x141 // Unchanged 5.0
	InitZoneOpcode      = 0x19B // Updated 5.1
	ControlOpcode       = 0x264 // Updated 5.1
	ControlSelfOpcode   = 0x164 // Updated 5.1
	ControlTargetOpcode = 0x16C // Updated 5.1
	RemoveEntityOpcode  = 0x097 // Updated 5.1
	UpdateHPMPTPOpcode  = 0x32D // Updated 5.1

	UpdateStatusesOpcode       = 0x15B // Updated for 5.0
	UpdateStatusesEurekaOpcode = 0x15C // Updated for 5.0

	ActionOpcode      = 0xA7  // Updated for 5.1
	AoEAction8Opcode  = 0xA9  // Updated for 5.1
	AoEAction16Opcode = 0x15F // Updated for 5.1
	AoEAction24Opcode = 0x292 // Updated for 5.1
	AoEAction32Opcode = 0x268 // Updated for 5.1

	PlayerSpawnOpcode = 0x386 // Updated for 5.1
	NPCSpawnOpcode    = 0x10A // Updated for 5.1
	NPCSpawn2Opcode   = 0x115 // Updated for 5.1

	MovementOpcode = 0x1BC // Updated for 5.1
	SetPosOpcode   = 0x311 // Updated for 5.1

	CastingOpcode = 0x12C // Updated for 5.1

	HateRankingOpcode = 0x354 // Updated for 5.1
	HateListOpcode    = 0x0C7 // Updated for 5.1

	EquipChangeOpcode = 0x25E // Updated for 5.1

	EventPlayOpcode         = 0x1B5 // Updated for 5.0
	EventPlay2Opcode        = 0x1B6 // Updated for 5.0
	DirectorPlaySceneOpcode = 0x1B9 // Updated for 5.0

	MountOpcode = 0x1F3 // Updated for 5.0

	WeatherChangeOpcode = 0x210 // Updated for 5.0

	WaymarkOpcode = 0x272 // Updated for 5.0

	PrepareZoningOpcode = 0x2A4 // Updated for 5.0

	GaugeOpcode = 0x2A5 // Updated for 5.0
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
	registerInBlockFactory(UpdateStatusesEurekaOpcode, func() xivnet.BlockData { return new(UpdateStatusesEureka) })

	registerInBlockFactory(ActionOpcode, func() xivnet.BlockData { return new(Action) })
	registerInBlockFactory(AoEAction8Opcode, func() xivnet.BlockData { return new(AoEAction8) })
	registerInBlockFactory(AoEAction16Opcode, func() xivnet.BlockData { return new(AoEAction16) })
	registerInBlockFactory(AoEAction24Opcode, func() xivnet.BlockData { return new(AoEAction24) })
	registerInBlockFactory(AoEAction32Opcode, func() xivnet.BlockData { return new(AoEAction32) })

	registerInBlockFactory(PlayerSpawnOpcode, func() xivnet.BlockData { return new(PlayerSpawn) })
	registerInBlockFactory(NPCSpawnOpcode, func() xivnet.BlockData { return new(NPCSpawn) })
	registerInBlockFactory(NPCSpawn2Opcode, func() xivnet.BlockData { return new(NPCSpawn2) })

	registerInBlockFactory(MovementOpcode, func() xivnet.BlockData { return new(Movement) })
	registerInBlockFactory(SetPosOpcode, func() xivnet.BlockData { return new(SetPos) })
	registerInBlockFactory(CastingOpcode, func() xivnet.BlockData { return new(Casting) })

	registerInBlockFactory(HateRankingOpcode, func() xivnet.BlockData { return new(HateRanking) })
	registerInBlockFactory(HateListOpcode, func() xivnet.BlockData { return new(HateList) })

	registerInBlockFactory(EquipChangeOpcode, func() xivnet.BlockData { return new(EquipChange) })

	registerInBlockFactory(EventPlayOpcode, func() xivnet.BlockData { return new(EventPlay) })
	registerInBlockFactory(EventPlay2Opcode, func() xivnet.BlockData { return new(EventPlay2) })
	registerInBlockFactory(DirectorPlaySceneOpcode, func() xivnet.BlockData { return new(DirectorPlayScene) })

	registerInBlockFactory(MountOpcode, func() xivnet.BlockData { return new(Mount) })

	registerInBlockFactory(WeatherChangeOpcode, func() xivnet.BlockData { return new(WeatherChange) })

	registerInBlockFactory(WaymarkOpcode, func() xivnet.BlockData { return new(Marker) })
	registerInBlockFactory(PrepareZoningOpcode, func() xivnet.BlockData { return new(PrepareZoning) })

	registerInBlockFactory(GaugeOpcode, func() xivnet.BlockData { return new(Gauge) })
	// registerInBlockFactory(PerformOpcode, func() xivnet.BlockData { return new(Perform) })

	// registerInBlockFactory(XWorldPartyListOpcode, new(XWorldPartyList))
}

// Opcodes that define the datatypes of outgoing (to server) network blocks
const (
	EgressClientTriggerOpcode = 0x13A // Updated for 5.0

	EgressMovementOpcode         = 0x141 // Updated for 5.0
	EgressInstanceMovementOpcode = 0x180 // Updated for 5.0

	EgressPerformOpcode    = 0x18B // Updated for 5.0
	EgressCraftEventOpcode = 0x15F // Updated for 5.0
)

func init() {
	registerOutBlockFactory(EgressClientTriggerOpcode, func() xivnet.BlockData { return new(EgressClientTrigger) })

	registerOutBlockFactory(EgressMovementOpcode, func() xivnet.BlockData { return new(EgressMovement) })
	registerOutBlockFactory(EgressInstanceMovementOpcode, func() xivnet.BlockData { return new(EgressInstanceMovement) })

	registerOutBlockFactory(EgressPerformOpcode, func() xivnet.BlockData { return new(Perform) })

	registerOutBlockFactory(EgressCraftEventOpcode, func() xivnet.BlockData { return new(EgressCraftEvent) })
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
