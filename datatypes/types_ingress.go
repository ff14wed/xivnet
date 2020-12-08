package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x279 // Updated for 5.4
	InitZoneOpcode      = 0x3D0 // Updated for 5.4
	ControlOpcode       = 0x117 // Updated for 5.4
	ControlSelfOpcode   = 0x318 // Updated for 5.4
	ControlTargetOpcode = 0x2C6 // Updated for 5.4
	RemoveEntityOpcode  = 0x24B // Updated for 5.4
	UpdateHPMPTPOpcode  = 0x3B5 // Updated for 5.4

	ChatZoneOpcode = 0x3D8 // Updated for 5.4

	UpdateStatusesOpcode       = 0x177 // Updated for 5.4
	UpdateStatusesEurekaOpcode = 0x30D // Updated for 5.4
	UpdateStatusesBossOpcode   = 0x1DD // Updated for 5.4

	ActionOpcode      = 0x357 // Updated for 5.4
	AoEAction8Opcode  = 0x3BD // Updated for 5.4
	AoEAction16Opcode = 0x12F // Updated for 5.4
	AoEAction24Opcode = 0x384 // Updated for 5.4
	AoEAction32Opcode = 0x3A9 // Updated for 5.4

	ObjectSpawnOpcode = 0x15F // Updated for 5.4
	PlayerSpawnOpcode = 0x33A // Updated for 5.4
	NPCSpawnOpcode    = 0x365 // Updated for 5.4
	NPCSpawn2Opcode   = 0x137 // Updated for 5.4

	MovementOpcode = 0x293 // Updated for 5.4
	SetPosOpcode   = 0x341 // Updated for 5.4

	CastingOpcode = 0x25D // Updated for 5.4

	HateRankingOpcode = 0x2B2 // Updated for 5.4
	HateListOpcode    = 0x107 // Updated for 5.4

	EquipChangeOpcode = 0x93 // Updated for 5.4

	EventPlayOpcode    = 0x378 // Updated for 5.4
	EventPlay4Opcode   = 0x145 // Updated for 5.4
	EventPlay8Opcode   = 0x99  // Updated for 5.4
	EventPlay16Opcode  = 0xD0  // Updated for 5.4
	EventPlay32Opcode  = 0x18E // Updated for 5.4
	EventPlay64Opcode  = 0x248 // Updated for 5.4
	EventPlay128Opcode = 0x265 // Updated for 5.4
	EventPlay255Opcode = 0x273 // Updated for 5.4

	MountOpcode = 0xCD // Updated for 5.4

	WeatherChangeOpcode = 0x2EF // Updated for 5.4

	PrepareZoningOpcode = 0x26E // Updated for 5.4

	GaugeOpcode = 0x1ED // Updated for 5.4

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
