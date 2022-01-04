package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x203 // Updated for 6.05
	InitZoneOpcode      = 0x137 // Updated for 6.05
	ControlOpcode       = 0x2CF // Updated for 6.05
	ControlSelfOpcode   = 0x96  // Updated for 6.05
	ControlTargetOpcode = 0x272 // Updated for 6.05
	RemoveEntityOpcode  = 0x227 // Updated for 6.05
	UpdateHPMPTPOpcode  = 0xF4  // Updated for 6.05

	ChatZoneOpcode = 0x32F // Updated for 6.05

	UpdateStatusesOpcode       = 0x188 // Updated for 6.05
	UpdateStatusesEurekaOpcode = 0xBA  // Updated for 6.05
	UpdateStatusesBossOpcode   = 0x38F // Updated for 6.05

	ActionOpcode      = 0x33E // Updated for 6.05
	AoEAction8Opcode  = 0x1F4 // Updated for 6.05
	AoEAction16Opcode = 0x1FA // Updated for 6.05
	AoEAction24Opcode = 0x300 // Updated for 6.05
	AoEAction32Opcode = 0x3CD // Updated for 6.05

	ObjectSpawnOpcode = 0x1FD // Updated for 6.05
	PlayerSpawnOpcode = 0x338 // Updated for 6.05
	NPCSpawnOpcode    = 0x1D2 // Updated for 6.05
	NPCSpawn2Opcode   = 0x270 // Updated for 6.05

	MovementOpcode = 0xDB // Updated for 6.05
	SetPosOpcode   = 0x81 // Updated for 6.05

	CastingOpcode = 0x307 // Updated for 6.05

	HateRankingOpcode = 0x3A1 // Updated for 6.05
	HateListOpcode    = 0x26E // Updated for 6.05

	EquipChangeOpcode = 0x2EE // Updated for 6.05

	EventPlayOpcode    = 0x13F // Updated for 6.05
	EventPlay4Opcode   = 0x212 // Updated for 6.05
	EventPlay8Opcode   = 0x10B // Updated for 6.05
	EventPlay16Opcode  = 0xD0  // Updated for 6.05
	EventPlay32Opcode  = 0xC5  // Updated for 6.05
	EventPlay64Opcode  = 0xC6  // Updated for 6.05
	EventPlay128Opcode = 0x32C // Updated for 6.05
	EventPlay255Opcode = 0x295 // Updated for 6.05

	MountOpcode = 0x26F // Updated for 6.05

	WeatherChangeOpcode = 0x1EE // Updated for 6.05

	PrepareZoningOpcode = 0x1DD // Updated for 6.05

	GaugeOpcode = 0x22D // Updated for 6.05

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
