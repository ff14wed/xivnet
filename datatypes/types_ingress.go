package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x37B // Updated for 6.35
	InitZoneOpcode      = 0x37D // Updated for 6.35
	ControlOpcode       = 0x1BB // Updated for 6.35
	ControlSelfOpcode   = 0x228 // Updated for 6.35
	ControlTargetOpcode = 0x19D // Updated for 6.35
	RemoveEntityOpcode  = 0xCD  // Updated for 6.35
	UpdateHPMPTPOpcode  = 0x394 // Updated for 6.35

	ChatZoneOpcode = 0x118 // Updated for 6.35

	UpdateStatusesOpcode       = 0x2D4 // Updated for 6.35
	UpdateStatusesEurekaOpcode = 0x3D0 // Updated for 6.35
	UpdateStatusesBossOpcode   = 0x35D // Updated for 6.35

	ActionOpcode      = 0xD4  // Updated for 6.35
	AoEAction8Opcode  = 0x1A4 // Updated for 6.35
	AoEAction16Opcode = 0x1C9 // Updated for 6.35
	AoEAction24Opcode = 0x252 // Updated for 6.35
	AoEAction32Opcode = 0x2C8 // Updated for 6.35

	ObjectSpawnOpcode = 0x1F4 // Updated for 6.35
	PlayerSpawnOpcode = 0x100 // Updated for 6.35
	NPCSpawnOpcode    = 0x1C0 // Updated for 6.35
	NPCSpawn2Opcode   = 0x163 // Updated for 6.35

	MovementOpcode = 0xB4  // Updated for 6.35
	SetPosOpcode   = 0x2E5 // Updated for 6.35

	CastingOpcode = 0x291 // Updated for 6.35

	HateRankingOpcode = 0x17A // Updated for 6.35
	HateListOpcode    = 0x205 // Updated for 6.35

	PlayerStatsOpcode = 0xEE // Updated for 6.35

	EquipChangeOpcode = 0x381 // Updated for 6.35

	EventPlayOpcode    = 0x17F // Updated for 6.35
	EventPlay4Opcode   = 0xCA  // Updated for 6.35
	EventPlay8Opcode   = 0x175 // Updated for 6.35
	EventPlay16Opcode  = 0x26A // Updated for 6.35
	EventPlay32Opcode  = 0x74  // Updated for 6.35
	EventPlay64Opcode  = 0x255 // Updated for 6.35
	EventPlay128Opcode = 0x330 // Updated for 6.35
	EventPlay255Opcode = 0xDD  // Updated for 6.35

	MountOpcode = 0x317 // Updated for 6.35

	WeatherChangeOpcode = 0x371 // Updated for 6.35

	PrepareZoningOpcode = 0x99 // Updated for 6.35

	GaugeOpcode = 0x376 // Updated for 6.35

	WaymarkOpcode         = 0x160 // Updated for 6.35
	PerformOpcode         = 0xC6  // Updated for 6.35
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
