package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x24A // Updated for 6.15
	InitZoneOpcode      = 0x184 // Updated for 6.15
	ControlOpcode       = 0x286 // Updated for 6.15
	ControlSelfOpcode   = 0x308 // Updated for 6.15
	ControlTargetOpcode = 0x70  // Updated for 6.15
	RemoveEntityOpcode  = 0x3C6 // Updated for 6.15
	UpdateHPMPTPOpcode  = 0x2E1 // Updated for 6.15

	ChatZoneOpcode = 0x1EC // Updated for 6.15

	UpdateStatusesOpcode       = 0x3B9 // Updated for 6.15
	UpdateStatusesEurekaOpcode = 0x16A // Updated for 6.15
	UpdateStatusesBossOpcode   = 0x128 // Updated for 6.15

	ActionOpcode      = 0x14C // Updated for 6.15
	AoEAction8Opcode  = 0xF4  // Updated for 6.15
	AoEAction16Opcode = 0x2B0 // Updated for 6.15
	AoEAction24Opcode = 0x2F9 // Updated for 6.15
	AoEAction32Opcode = 0x15B // Updated for 6.15

	ObjectSpawnOpcode = 0x1B7 // Updated for 6.15
	PlayerSpawnOpcode = 0x2DD // Updated for 6.15
	NPCSpawnOpcode    = 0x331 // Updated for 6.15
	NPCSpawn2Opcode   = 0x376 // Updated for 6.15

	MovementOpcode = 0x1F1 // Updated for 6.15
	SetPosOpcode   = 0x3C8 // Updated for 6.15

	CastingOpcode = 0x347 // Updated for 6.15

	HateRankingOpcode = 0xD7 // Updated for 6.15
	HateListOpcode    = 0xDC // Updated for 6.15

	PlayerStatsOpcode = 0x366 // Updated for 6.15

	EquipChangeOpcode = 0x193 // Updated for 6.15

	EventPlayOpcode    = 0x72  // Updated for 6.15
	EventPlay4Opcode   = 0x1AF // Updated for 6.15
	EventPlay8Opcode   = 0x35D // Updated for 6.15
	EventPlay16Opcode  = 0x374 // Updated for 6.15
	EventPlay32Opcode  = 0x34F // Updated for 6.15
	EventPlay64Opcode  = 0x92  // Updated for 6.15
	EventPlay128Opcode = 0xB3  // Updated for 6.15

	MountOpcode = 0x218 // Updated for 6.15

	WeatherChangeOpcode = 0x6C // Updated for 6.15

	PrepareZoningOpcode = 0x1D0 // Updated for 6.15

	GaugeOpcode = 0x1FC // Updated for 6.15

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
