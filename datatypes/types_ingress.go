package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x3B5 // Updated for 6.38
	InitZoneOpcode      = 0x1FE // Updated for 6.38
	ControlOpcode       = 0x2C2 // Updated for 6.38
	ControlSelfOpcode   = 0x256 // Updated for 6.38
	ControlTargetOpcode = 0x1B8 // Updated for 6.38
	RemoveEntityOpcode  = 0x3A2 // Updated for 6.38
	UpdateHPMPTPOpcode  = 0x268 // Updated for 6.38

	ChatZoneOpcode = 0x1E4 // Updated for 6.38

	UpdateStatusesOpcode       = 0x317 // Updated for 6.38
	UpdateStatusesEurekaOpcode = 0x23B // Updated for 6.38
	UpdateStatusesBossOpcode   = 0xC8  // Updated for 6.38

	ActionOpcode      = 0xFB  // Updated for 6.38
	AoEAction8Opcode  = 0x2EF // Updated for 6.38
	AoEAction16Opcode = 0x3C6 // Updated for 6.38
	AoEAction24Opcode = 0x97  // Updated for 6.38
	AoEAction32Opcode = 0x24B // Updated for 6.38

	ObjectSpawnOpcode = 0x2D1 // Updated for 6.38
	PlayerSpawnOpcode = 0x94  // Updated for 6.38
	NPCSpawnOpcode    = 0x269 // Updated for 6.38
	NPCSpawn2Opcode   = 0x196 // Updated for 6.38

	MovementOpcode = 0x25B // Updated for 6.38
	SetPosOpcode   = 0x2F0 // Updated for 6.38

	CastingOpcode = 0x15F // Updated for 6.38

	HateRankingOpcode = 0xE5 // Updated for 6.38
	HateListOpcode    = 0x99 // Updated for 6.38

	PlayerStatsOpcode = 0x2D8 // Updated for 6.38

	EquipChangeOpcode = 0x2E2 // Updated for 6.38

	EventPlayOpcode    = 0x284 // Updated for 6.38
	EventPlay4Opcode   = 0xE2  // Updated for 6.38
	EventPlay8Opcode   = 0x2C1 // Updated for 6.38
	EventPlay16Opcode  = 0x7F  // Updated for 6.38
	EventPlay32Opcode  = 0xD3  // Updated for 6.38
	EventPlay64Opcode  = 0xC1  // Updated for 6.38
	EventPlay128Opcode = 0x29D // Updated for 6.38
	EventPlay255Opcode = 0x326 // Updated for 6.38

	MountOpcode = 0x2DF // Updated for 6.38

	WeatherChangeOpcode = 0x17A // Updated for 6.38

	PrepareZoningOpcode = 0x267 // Updated for 6.38

	GaugeOpcode = 0x136 // Updated for 6.38

	WaymarkOpcode         = 0x354 // Updated for 6.38
	PerformOpcode         = 0xB3  // Updated for 6.38
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
