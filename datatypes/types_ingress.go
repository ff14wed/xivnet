package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x93  // Updated for 5.55
	InitZoneOpcode      = 0x2D1 // Updated for 5.55
	ControlOpcode       = 0x30F // Updated for 5.55
	ControlSelfOpcode   = 0x200 // Updated for 5.55
	ControlTargetOpcode = 0x387 // Updated for 5.55
	RemoveEntityOpcode  = 0x233 // Updated for 5.55
	UpdateHPMPTPOpcode  = 0x1E8 // Updated for 5.55

	ChatZoneOpcode = 0xFF // Updated for 5.55

	UpdateStatusesOpcode       = 0x117 // Updated for 5.55
	UpdateStatusesEurekaOpcode = 0x372 // Updated for 5.55
	UpdateStatusesBossOpcode   = 0x3C7 // Updated for 5.55

	ActionOpcode      = 0x128 // Updated for 5.55
	AoEAction8Opcode  = 0x295 // Updated for 5.55
	AoEAction16Opcode = 0x25E // Updated for 5.55
	AoEAction24Opcode = 0x299 // Updated for 5.55
	AoEAction32Opcode = 0xA7  // Updated for 5.55

	ObjectSpawnOpcode = 0x207 // Updated for 5.55
	PlayerSpawnOpcode = 0x8B  // Updated for 5.55
	NPCSpawnOpcode    = 0xE0  // Updated for 5.55
	NPCSpawn2Opcode   = 0x17A // Updated for 5.55

	MovementOpcode = 0x122 // Updated for 5.55
	SetPosOpcode   = 0x271 // Updated for 5.55

	CastingOpcode = 0x228 // Updated for 5.55

	HateRankingOpcode = 0x1E6 // Updated for 5.55
	HateListOpcode    = 0x132 // Updated for 5.55

	EquipChangeOpcode = 0x31B // Updated for 5.55

	EventPlayOpcode    = 0x183 // Updated for 5.55
	EventPlay4Opcode   = 0x38B // Updated for 5.55
	EventPlay8Opcode   = 0x32D // Updated for 5.55
	EventPlay16Opcode  = 0x346 // Updated for 5.55
	EventPlay32Opcode  = 0x3A0 // Updated for 5.55
	EventPlay64Opcode  = 0x18D // Updated for 5.55
	EventPlay128Opcode = 0x290 // Updated for 5.55
	EventPlay255Opcode = 0x29C // Updated for 5.55

	MountOpcode = 0x33B // Updated for 5.55

	WeatherChangeOpcode = 0x2FF // Updated for 5.55

	PrepareZoningOpcode = 0x21B // Updated for 5.55

	GaugeOpcode = 0x138 // Updated for 5.55

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
