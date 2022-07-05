package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x95  // Updated for 6.18
	InitZoneOpcode      = 0x35D // Updated for 6.18
	ControlOpcode       = 0x36E // Updated for 6.18
	ControlSelfOpcode   = 0x2F0 // Updated for 6.18
	ControlTargetOpcode = 0x1FD // Updated for 6.18
	RemoveEntityOpcode  = 0x265 // Updated for 6.18
	UpdateHPMPTPOpcode  = 0x20F // Updated for 6.18

	ChatZoneOpcode = 0x166 // Updated for 6.18

	UpdateStatusesOpcode       = 0xE9  // Updated for 6.18
	UpdateStatusesEurekaOpcode = 0x396 // Updated for 6.18
	UpdateStatusesBossOpcode   = 0x267 // Updated for 6.18

	ActionOpcode      = 0xEE  // Updated for 6.18
	AoEAction8Opcode  = 0x3AA // Updated for 6.18
	AoEAction16Opcode = 0x2BA // Updated for 6.18
	AoEAction24Opcode = 0x211 // Updated for 6.18
	AoEAction32Opcode = 0x2D5 // Updated for 6.18

	ObjectSpawnOpcode = 0x117 // Updated for 6.18
	PlayerSpawnOpcode = 0x1D7 // Updated for 6.18
	NPCSpawnOpcode    = 0x86  // Updated for 6.18
	NPCSpawn2Opcode   = 0x3C0 // Updated for 6.18

	MovementOpcode = 0x188 // Updated for 6.18
	SetPosOpcode   = 0x2C4 // Updated for 6.18

	CastingOpcode = 0x3C2 // Updated for 6.18

	HateRankingOpcode = 0x3A5 // Updated for 6.18
	HateListOpcode    = 0x193 // Updated for 6.18

	PlayerStatsOpcode = 0x2AE // Updated for 6.18

	EquipChangeOpcode = 0x9C // Updated for 6.18

	EventPlayOpcode    = 0x348 // Updated for 6.18
	EventPlay4Opcode   = 0x108 // Updated for 6.18
	EventPlay8Opcode   = 0x10F // Updated for 6.18
	EventPlay16Opcode  = 0x2D3 // Updated for 6.18
	EventPlay32Opcode  = 0x2EF // Updated for 6.18
	EventPlay64Opcode  = 0x192 // Updated for 6.18
	EventPlay128Opcode = 0x155 // Updated for 6.18
	EventPlay255Opcode = 0x143 // Updated for 6.18

	MountOpcode = 0x236 // Updated for 6.18

	WeatherChangeOpcode = 0xDE // Updated for 6.18

	PrepareZoningOpcode = 0x346 // Updated for 6.18

	GaugeOpcode = 0x1E9 // Updated for 6.18

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
