package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x36C // Updated for 6.31
	InitZoneOpcode      = 0x94  // Updated for 6.31
	ControlOpcode       = 0x363 // Updated for 6.31
	ControlSelfOpcode   = 0x267 // Updated for 6.31
	ControlTargetOpcode = 0x1EC // Updated for 6.31
	RemoveEntityOpcode  = 0x1C3 // Updated for 6.31
	UpdateHPMPTPOpcode  = 0x10D // Updated for 6.31

	ChatZoneOpcode = 0x353 // Updated for 6.31

	UpdateStatusesOpcode       = 0x2A4 // Updated for 6.31
	UpdateStatusesEurekaOpcode = 0x1DE // Updated for 6.31
	UpdateStatusesBossOpcode   = 0xA6  // Updated for 6.31

	ActionOpcode      = 0x3C1 // Updated for 6.31
	AoEAction8Opcode  = 0x78  // Updated for 6.31
	AoEAction16Opcode = 0x398 // Updated for 6.31
	AoEAction24Opcode = 0x2EA // Updated for 6.31
	AoEAction32Opcode = 0x210 // Updated for 6.31

	ObjectSpawnOpcode = 0x11A // Updated for 6.31
	PlayerSpawnOpcode = 0x187 // Updated for 6.31
	NPCSpawnOpcode    = 0x391 // Updated for 6.31
	NPCSpawn2Opcode   = 0x225 // Updated for 6.31

	MovementOpcode = 0x2A1 // Updated for 6.31
	SetPosOpcode   = 0x186 // Updated for 6.31

	CastingOpcode = 0x207 // Updated for 6.31

	HateRankingOpcode = 0x250 // Updated for 6.31
	HateListOpcode    = 0x359 // Updated for 6.31

	PlayerStatsOpcode = 0x272 // Updated for 6.31

	EquipChangeOpcode = 0x212 // Updated for 6.31

	EventPlayOpcode    = 0x1F5 // Updated for 6.31
	EventPlay4Opcode   = 0x357 // Updated for 6.31
	EventPlay8Opcode   = 0x269 // Updated for 6.31
	EventPlay16Opcode  = 0x278 // Updated for 6.31
	EventPlay32Opcode  = 0x36B // Updated for 6.31
	EventPlay64Opcode  = 0x288 // Updated for 6.31
	EventPlay128Opcode = 0x73  // Updated for 6.31
	EventPlay255Opcode = 0x23A // Updated for 6.31

	MountOpcode = 0x16B // Updated for 6.31

	WeatherChangeOpcode = 0x163 // Updated for 6.31

	PrepareZoningOpcode = 0x1D7 // Updated for 6.31

	GaugeOpcode = 0xA9 // Updated for 6.31

	WaymarkOpcode         = 0x175 // Updated for 6.31
	PerformOpcode         = 0x2AB // Updated for 6.31
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
	registerInBlockFactory(EventPlay8Opcode, func() xivnet.BlockData { return new(EventPlay8) })
	registerInBlockFactory(EventPlay16Opcode, func() xivnet.BlockData { return new(EventPlay16) })
	registerInBlockFactory(EventPlay32Opcode, func() xivnet.BlockData { return new(EventPlay32) })
	registerInBlockFactory(EventPlay64Opcode, func() xivnet.BlockData { return new(EventPlay64) })
	registerInBlockFactory(EventPlay128Opcode, func() xivnet.BlockData { return new(EventPlay128) })
	registerInBlockFactory(EventPlay255Opcode, func() xivnet.BlockData { return new(EventPlay255) })

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
