package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x1A8 // Updated for 6.2
	InitZoneOpcode      = 0x142 // Updated for 6.2
	ControlOpcode       = 0x2B4 // Updated for 6.2
	ControlSelfOpcode   = 0x1AE // Updated for 6.2
	ControlTargetOpcode = 0xAC  // Updated for 6.2
	RemoveEntityOpcode  = 0x196 // Updated for 6.2
	UpdateHPMPTPOpcode  = 0xC7  // Updated for 6.2

	ChatZoneOpcode = 0x19E // Updated for 6.2

	UpdateStatusesOpcode       = 0x8D  // Updated for 6.2
	UpdateStatusesEurekaOpcode = 0x264 // Updated for 6.2
	UpdateStatusesBossOpcode   = 0x2A0 // Updated for 6.2

	ActionOpcode      = 0x350 // Updated for 6.2
	AoEAction8Opcode  = 0x230 // Updated for 6.2
	AoEAction16Opcode = 0x385 // Updated for 6.2
	AoEAction24Opcode = 0x3E2 // Updated for 6.2
	AoEAction32Opcode = 0x1EC // Updated for 6.2

	ObjectSpawnOpcode = 0x3AC // Updated for 6.2
	PlayerSpawnOpcode = 0x176 // Updated for 6.2
	NPCSpawnOpcode    = 0x1EF // Updated for 6.2
	NPCSpawn2Opcode   = 0x3DF // Updated for 6.2

	MovementOpcode = 0x10B // Updated for 6.2
	SetPosOpcode   = 0x1C0 // Updated for 6.2

	CastingOpcode = 0x3CC // Updated for 6.2

	HateRankingOpcode = 0x101 // Updated for 6.2
	HateListOpcode    = 0x147 // Updated for 6.2

	PlayerStatsOpcode = 0x313 // Updated for 6.2

	EquipChangeOpcode = 0x390 // Updated for 6.2

	EventPlayOpcode    = 0x66  // Updated for 6.2
	EventPlay4Opcode   = 0x233 // Updated for 6.2
	EventPlay8Opcode   = 0x1f8 // Updated for 6.2
	EventPlay16Opcode  = 0xfa  // Updated for 6.2
	EventPlay32Opcode  = 0x35b // Updated for 6.2
	EventPlay64Opcode  = 0xaa  // Updated for 6.2
	EventPlay128Opcode = 0x213 // Updated for 6.2
	EventPlay255Opcode = 0x298 // Updated for 6.2

	MountOpcode = 0x38E // Updated for 6.2

	WeatherChangeOpcode = 0x143 // Updated for 6.2

	PrepareZoningOpcode = 0x1E7 // Updated for 6.2

	GaugeOpcode = 0x394 // Updated for 6.2

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
