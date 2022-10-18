package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x2ca // Updated for 6.25
	InitZoneOpcode      = 0x144 // Updated for 6.25
	ControlOpcode       = 0x39E // Updated for 6.25
	ControlSelfOpcode   = 0x3CE // Updated for 6.25
	ControlTargetOpcode = 0x3C6 // Updated for 6.25
	RemoveEntityOpcode  = 0x271 // Updated for 6.25
	UpdateHPMPTPOpcode  = 0x129 // Updated for 6.25

	ChatZoneOpcode = 0x333 // Updated for 6.25

	UpdateStatusesOpcode       = 0x274 // Updated for 6.25
	UpdateStatusesEurekaOpcode = 0x194 // Updated for 6.25
	UpdateStatusesBossOpcode   = 0x2C3 // Updated for 6.25

	ActionOpcode      = 0x3C4 // Updated for 6.25
	AoEAction8Opcode  = 0x391 // Updated for 6.25
	AoEAction16Opcode = 0x2E4 // Updated for 6.25
	AoEAction24Opcode = 0x2EF // Updated for 6.25
	AoEAction32Opcode = 0x3B9 // Updated for 6.25

	ObjectSpawnOpcode = 0x2D9 // Updated for 6.25
	PlayerSpawnOpcode = 0x17E // Updated for 6.25
	NPCSpawnOpcode    = 0x29D // Updated for 6.25
	NPCSpawn2Opcode   = 0x38B // Updated for 6.25

	MovementOpcode = 0x275 // Updated for 6.25
	SetPosOpcode   = 0x28F // Updated for 6.25

	CastingOpcode = 0x2EC // Updated for 6.25

	HateRankingOpcode = 0x221 // Updated for 6.25
	HateListOpcode    = 0x112 // Updated for 6.25

	PlayerStatsOpcode = 0x1C3 // Updated for 6.25

	EquipChangeOpcode = 0x39C // Updated for 6.25

	EventPlayOpcode    = 0x162 // Updated for 6.25
	EventPlay4Opcode   = 0x232 // Updated for 6.25
	EventPlay8Opcode   = 0xe3  // Updated for 6.25
	EventPlay16Opcode  = 0x29e // Updated for 6.25
	EventPlay32Opcode  = 0x241 // Updated for 6.25
	EventPlay64Opcode  = 0x3a1 // Updated for 6.25
	EventPlay128Opcode = 0x38c // Updated for 6.25
	EventPlay255Opcode = 0x257 // Updated for 6.25

	MountOpcode = 0x242 // Updated for 6.25

	WeatherChangeOpcode = 0x352 // Updated for 6.25

	PrepareZoningOpcode = 0x1AC // Updated for 6.25

	GaugeOpcode = 0x344 // Updated for 6.25

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

	registerInBlockFactory(PlayerStatsOpcode, func() xivnet.BlockData { return new(PlayerStats) })

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
