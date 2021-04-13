package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x143 // Updated for 5.5
	InitZoneOpcode      = 0x3D7 // Updated for 5.5
	ControlOpcode       = 0x3B5 // Updated for 5.5
	ControlSelfOpcode   = 0x3C0 // Updated for 5.5
	ControlTargetOpcode = 0x202 // Updated for 5.5
	RemoveEntityOpcode  = 0xC2  // Updated for 5.5
	UpdateHPMPTPOpcode  = 0xEF  // Updated for 5.5

	ChatZoneOpcode = 0x2B0 // Updated for 5.5

	UpdateStatusesOpcode       = 0x343 // Updated for 5.5
	UpdateStatusesEurekaOpcode = 0x289 // Updated for 5.5
	UpdateStatusesBossOpcode   = 0x34D // Updated for 5.5

	ActionOpcode      = 0x204 // Updated for 5.5
	AoEAction8Opcode  = 0x1C7 // Updated for 5.5
	AoEAction16Opcode = 0x32A // Updated for 5.5
	AoEAction24Opcode = 0x35C // Updated for 5.5
	AoEAction32Opcode = 0x2C5 // Updated for 5.5

	ObjectSpawnOpcode = 0x358 // Updated for 5.5
	PlayerSpawnOpcode = 0x2EF // Updated for 5.5
	NPCSpawnOpcode    = 0x3C1 // Updated for 5.5
	NPCSpawn2Opcode   = 0x82  // Updated for 5.5

	MovementOpcode = 0x3AE // Updated for 5.5
	SetPosOpcode   = 0x342 // Updated for 5.5

	CastingOpcode = 0xB6 // Updated for 5.5

	HateRankingOpcode = 0xCA // Updated for 5.5
	HateListOpcode    = 0x89 // Updated for 5.5

	EquipChangeOpcode = 0x36F // Updated for 5.5

	EventPlayOpcode    = 0x15B // Updated for 5.5
	EventPlay4Opcode   = 0x7E  // Updated for 5.5
	EventPlay8Opcode   = 0x124 // Updated for 5.5
	EventPlay16Opcode  = 0x200 // Updated for 5.5
	EventPlay32Opcode  = 0x2CE // Updated for 5.5
	EventPlay64Opcode  = 0x1FA // Updated for 5.5
	EventPlay128Opcode = 0x308 // Updated for 5.5
	EventPlay255Opcode = 0x236 // Updated for 5.5

	MountOpcode = 0x100 // Updated for 5.5

	WeatherChangeOpcode = 0x22F // Updated for 5.5

	PrepareZoningOpcode = 0xD2 // Updated for 5.5

	GaugeOpcode = 0x3A7 // Updated for 5.5

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
