package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x1C2 // Updated for 5.35a
	InitZoneOpcode      = 0x3CD // Updated for 5.35a
	ControlOpcode       = 0x2A4 // Updated for 5.35a
	ControlSelfOpcode   = 0x2C8 // Updated for 5.35a
	ControlTargetOpcode = 0x209 // Updated for 5.35a
	RemoveEntityOpcode  = 0x239 // Updated for 5.35a
	UpdateHPMPTPOpcode  = 0x319 // Updated for 5.35a

	ChatZoneOpcode = 0x349 // Updated for 5.35a

	UpdateStatusesOpcode       = 0x382 // Updated for 5.35a
	UpdateStatusesEurekaOpcode = 0x342 // Updated for 5.35a
	UpdateStatusesBossOpcode   = 0x298 // Updated for 5.35a

	ActionOpcode      = 0x192 // Updated for 5.35a
	AoEAction8Opcode  = 0x12C // Updated for 5.35a
	AoEAction16Opcode = 0x1B9 // Updated for 5.35a
	AoEAction24Opcode = 0x2B4 // Updated for 5.35a
	AoEAction32Opcode = 0xA4  // Updated for 5.35a

	PlayerSpawnOpcode = 0x179 // Updated for 5.35a
	NPCSpawnOpcode    = 0x3A8 // Updated for 5.35a
	NPCSpawn2Opcode   = 0x26A // Updated for 5.35a

	MovementOpcode = 0x1BF // Updated for 5.35a
	SetPosOpcode   = 0x3DF // Updated for 5.35a

	CastingOpcode = 0x302 // Updated for 5.35a

	HateRankingOpcode = 0x2CC // Updated for 5.35a
	HateListOpcode    = 0x198 // Updated for 5.35a

	EquipChangeOpcode = 0x277 // Updated for 5.35a

	EventPlayOpcode   = 0xF3  // Updated for 5.35a
	EventPlay4Opcode  = 0xAC  // Updated for 5.35a
	EventPlay32Opcode = 0x29A // Updated for 5.35a

	MountOpcode = 0x1B5 // Updated for 5.35a

	WeatherChangeOpcode = 0x27B // Updated for 5.35a

	PrepareZoningOpcode = 0x26C // Updated for 5.35a

	GaugeOpcode = 0x112 // Updated for 5.35a

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
