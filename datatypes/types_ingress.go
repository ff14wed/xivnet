package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x151 // Updated for 5.58
	InitZoneOpcode      = 0x100 // Updated for 5.58
	ControlOpcode       = 0x264 // Updated for 5.58
	ControlSelfOpcode   = 0x314 // Updated for 5.58
	ControlTargetOpcode = 0xFC  // Updated for 5.58
	RemoveEntityOpcode  = 0x210 // Updated for 5.58
	UpdateHPMPTPOpcode  = 0x39B // Updated for 5.58

	ChatZoneOpcode = 0x220 // Updated for 5.58

	UpdateStatusesOpcode       = 0x1C5 // Updated for 5.58
	UpdateStatusesEurekaOpcode = 0x14F // Updated for 5.58
	UpdateStatusesBossOpcode   = 0x1C8 // Updated for 5.58

	ActionOpcode      = 0x102 // Updated for 5.58
	AoEAction8Opcode  = 0x345 // Updated for 5.58
	AoEAction16Opcode = 0x2B6 // Updated for 5.58
	AoEAction24Opcode = 0x298 // Updated for 5.58
	AoEAction32Opcode = 0x3A4 // Updated for 5.58

	ObjectSpawnOpcode = 0x104 // Updated for 5.58
	PlayerSpawnOpcode = 0x249 // Updated for 5.58
	NPCSpawnOpcode    = 0x14B // Updated for 5.58
	NPCSpawn2Opcode   = 0x1EB // Updated for 5.58

	MovementOpcode = 0x23D // Updated for 5.58
	SetPosOpcode   = 0x280 // Updated for 5.58

	CastingOpcode = 0x2A7 // Updated for 5.58

	HateRankingOpcode = 0x2C0 // Updated for 5.58
	HateListOpcode    = 0x1B4 // Updated for 5.58

	EquipChangeOpcode = 0x312 // Updated for 5.58

	EventPlayOpcode    = 0x1EF // Updated for 5.58
	EventPlay4Opcode   = 0x21C // Updated for 5.58
	EventPlay8Opcode   = 0x337 // Updated for 5.58
	EventPlay16Opcode  = 0x319 // Updated for 5.58
	EventPlay32Opcode  = 0x1E2 // Updated for 5.58
	EventPlay64Opcode  = 0x2FD // Updated for 5.58
	EventPlay128Opcode = 0x26E // Updated for 5.58
	EventPlay255Opcode = 0x39E // Updated for 5.58

	MountOpcode = 0x3C2 // Updated for 5.58

	WeatherChangeOpcode = 0x1B1 // Updated for 5.58

	PrepareZoningOpcode = 0x171 // Updated for 5.58

	GaugeOpcode = 0x335 // Updated for 5.58

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
