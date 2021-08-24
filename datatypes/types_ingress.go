package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x387 // Updated for 5.58a
	InitZoneOpcode      = 0x320 // Updated for 5.58a
	ControlOpcode       = 0xB0  // Updated for 5.58a
	ControlSelfOpcode   = 0x2B6 // Updated for 5.58a
	ControlTargetOpcode = 0x3C5 // Updated for 5.58a
	RemoveEntityOpcode  = 0xB5  // Updated for 5.58a
	UpdateHPMPTPOpcode  = 0x1A7 // Updated for 5.58a

	ChatZoneOpcode = 0xFE // Updated for 5.58a

	UpdateStatusesOpcode       = 0x74  // Updated for 5.58a
	UpdateStatusesEurekaOpcode = 0x19F // Updated for 5.58a
	UpdateStatusesBossOpcode   = 0x223 // Updated for 5.58a

	ActionOpcode      = 0x3CA // Updated for 5.58a
	AoEAction8Opcode  = 0x3C4 // Updated for 5.58a
	AoEAction16Opcode = 0xFA  // Updated for 5.58a
	AoEAction24Opcode = 0x339 // Updated for 5.58a
	AoEAction32Opcode = 0x23C // Updated for 5.58a

	ObjectSpawnOpcode = 0x125 // Updated for 5.58a
	PlayerSpawnOpcode = 0x1D8 // Updated for 5.58a
	NPCSpawnOpcode    = 0xD2  // Updated for 5.58a
	NPCSpawn2Opcode   = 0x18A // Updated for 5.58a

	MovementOpcode = 0xF8  // Updated for 5.58a
	SetPosOpcode   = 0x299 // Updated for 5.58a

	CastingOpcode = 0x15D // Updated for 5.58a

	HateRankingOpcode = 0x150 // Updated for 5.58a
	HateListOpcode    = 0x243 // Updated for 5.58a

	EquipChangeOpcode = 0x3A2 // Updated for 5.58a

	EventPlayOpcode    = 0x16B // Updated for 5.58a
	EventPlay4Opcode   = 0x10A // Updated for 5.58a
	EventPlay8Opcode   = 0x337 // Updated for 5.58a
	EventPlay16Opcode  = 0x269 // Updated for 5.58a
	EventPlay32Opcode  = 0x23E // Updated for 5.58a
	EventPlay64Opcode  = 0xDE  // Updated for 5.58a
	EventPlay128Opcode = 0x2D0 // Updated for 5.58a
	EventPlay255Opcode = 0x362 // Updated for 5.58a

	MountOpcode = 0x1E1 // Updated for 5.58a

	WeatherChangeOpcode = 0x323 // Updated for 5.58a

	PrepareZoningOpcode = 0x2AB // Updated for 5.58a

	GaugeOpcode = 0x1C1 // Updated for 5.58a

	WaymarkOpcode         = UndefinedOpcode
	PerformOpcode         = UndefinedOpcode
	XWorldPartyListOpcode = UndefinedOpcode
)

func init() {
	registerInBlockFactory(EffectResultOpcode, func() xivnet.BlockData { return new(EffectResult) })
	registerInBlockFactory(InitZoneOpcode, func() xivnet.BlockData { return new(InitZone) })
	registerInBlockFactory(ControlOpcode, func() xivnet.BlockData { return new(Control) })
	registerInBlockFactory(ControlSelfOpcode, func() xivnet.BlockData { return new(ControlSelf) })
	registerInBlockFactory(ControlTargetOpcode, func() xivnet.BlockData { return new(ControlTarget) })
	registerInBlockFactory(RemoveEntityOpcode, func() xivnet.BlockData { return new(RemoveEntity) })
	registerInBlockFactory(UpdateHPMPTPOpcode, func() xivnet.BlockData { return new(UpdateHPMPTP) })

	registerInBlockFactory(ChatZoneOpcode, func() xivnet.BlockData { return new(ChatZone) })

	registerInBlockFactory(UpdateStatusesOpcode, func() xivnet.BlockData { return new(UpdateStatuses) })
	registerInBlockFactory(UpdateStatusesEurekaOpcode, func() xivnet.BlockData { return new(UpdateStatusesEureka) })
	registerInBlockFactory(UpdateStatusesBossOpcode, func() xivnet.BlockData { return new(UpdateStatusesBoss) })

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
	registerInBlockFactory(EventPlay4Opcode, func() xivnet.BlockData { return new(EventPlay4) })
	registerInBlockFactory(EventPlay32Opcode, func() xivnet.BlockData { return new(EventPlay32) })

	registerInBlockFactory(MountOpcode, func() xivnet.BlockData { return new(Mount) })

	registerInBlockFactory(WeatherChangeOpcode, func() xivnet.BlockData { return new(WeatherChange) })

	registerInBlockFactory(WaymarkOpcode, func() xivnet.BlockData { return new(Marker) })
	registerInBlockFactory(PrepareZoningOpcode, func() xivnet.BlockData { return new(PrepareZoning) })

	registerInBlockFactory(GaugeOpcode, func() xivnet.BlockData { return new(Gauge) })
	registerInBlockFactory(PerformOpcode, func() xivnet.BlockData { return new(Perform) })

	registerInBlockFactory(XWorldPartyListOpcode, func() xivnet.BlockData { return new(XWorldPartyList) })
}

func registerInBlockFactory(opcode uint16, factory func() xivnet.BlockData) {
	inTypeRegistry[opcode] = factory
}
