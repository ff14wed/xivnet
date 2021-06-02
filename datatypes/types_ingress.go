package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x1E9 // Updated for 5.55a
	InitZoneOpcode      = 0x21C // Updated for 5.55a
	ControlOpcode       = 0x1C8 // Updated for 5.55a
	ControlSelfOpcode   = 0x35A // Updated for 5.55a
	ControlTargetOpcode = 0x338 // Updated for 5.55a
	RemoveEntityOpcode  = 0x1C5 // Updated for 5.55a
	UpdateHPMPTPOpcode  = 0x77  // Updated for 5.55a

	ChatZoneOpcode = 0x384 // Updated for 5.55a

	UpdateStatusesOpcode       = 0x18A // Updated for 5.55a
	UpdateStatusesEurekaOpcode = 0x36A // Updated for 5.55a
	UpdateStatusesBossOpcode   = 0x374 // Updated for 5.55a

	ActionOpcode      = 0x283 // Updated for 5.55a
	AoEAction8Opcode  = 0x25B // Updated for 5.55a
	AoEAction16Opcode = 0x15D // Updated for 5.55a
	AoEAction24Opcode = 0x91  // Updated for 5.55a
	AoEAction32Opcode = 0x169 // Updated for 5.55a

	ObjectSpawnOpcode = 0xFE  // Updated for 5.55a
	PlayerSpawnOpcode = 0x2C1 // Updated for 5.55a
	NPCSpawnOpcode    = 0xF1  // Updated for 5.55a
	NPCSpawn2Opcode   = 0x10A // Updated for 5.55a

	MovementOpcode = 0x9D  // Updated for 5.55a
	SetPosOpcode   = 0x266 // Updated for 5.55a

	CastingOpcode = 0xA9 // Updated for 5.55a

	HateRankingOpcode = 0x6E  // Updated for 5.55a
	HateListOpcode    = 0x1F1 // Updated for 5.55a

	EquipChangeOpcode = 0x264 // Updated for 5.55a

	EventPlayOpcode    = 0x369 // Updated for 5.55a
	EventPlay4Opcode   = 0x247 // Updated for 5.55a
	EventPlay8Opcode   = 0x227 // Updated for 5.55a
	EventPlay16Opcode  = 0xC5  // Updated for 5.55a
	EventPlay32Opcode  = 0x184 // Updated for 5.55a
	EventPlay64Opcode  = 0x11A // Updated for 5.55a
	EventPlay128Opcode = 0x1B3 // Updated for 5.55a
	EventPlay255Opcode = 0x240 // Updated for 5.55a

	MountOpcode = 0x16B // Updated for 5.55a

	WeatherChangeOpcode = 0x386 // Updated for 5.55a

	PrepareZoningOpcode = 0xA4 // Updated for 5.55a

	GaugeOpcode = 0x3B1 // Updated for 5.55a

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
