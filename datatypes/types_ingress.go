package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x124 // Updated for 5.4a
	InitZoneOpcode      = 0x35A // Updated for 5.4a
	ControlOpcode       = 0x9C  // Updated for 5.4a
	ControlSelfOpcode   = 0x6A  // Updated for 5.4a
	ControlTargetOpcode = 0x39E // Updated for 5.4a
	RemoveEntityOpcode  = 0x34A // Updated for 5.4a
	UpdateHPMPTPOpcode  = 0x17D // Updated for 5.4a

	ChatZoneOpcode = 0x3DC // Updated for 5.4a

	UpdateStatusesOpcode       = 0x1D1 // Updated for 5.4a
	UpdateStatusesEurekaOpcode = 0x17E // Updated for 5.4a
	UpdateStatusesBossOpcode   = 0x1C1 // Updated for 5.4a

	ActionOpcode      = 0x38F // Updated for 5.4a
	AoEAction8Opcode  = 0x21E // Updated for 5.4a
	AoEAction16Opcode = 0x248 // Updated for 5.4a
	AoEAction24Opcode = 0x82  // Updated for 5.4a
	AoEAction32Opcode = 0x1E3 // Updated for 5.4a

	ObjectSpawnOpcode = 0xE9  // Updated for 5.4a
	PlayerSpawnOpcode = 0x353 // Updated for 5.4a
	NPCSpawnOpcode    = 0x111 // Updated for 5.4a
	NPCSpawn2Opcode   = 0x199 // Updated for 5.4a

	MovementOpcode = 0x160 // Updated for 5.4a
	SetPosOpcode   = 0x25B // Updated for 5.4a

	CastingOpcode = 0x3CE // Updated for 5.4a

	HateRankingOpcode = 0x3AE // Updated for 5.4a
	HateListOpcode    = 0x20A // Updated for 5.4a

	EquipChangeOpcode = 0xCC // Updated for 5.4a

	EventPlayOpcode    = 0x336 // Updated for 5.4a
	EventPlay4Opcode   = 0x317 // Updated for 5.4a
	EventPlay8Opcode   = 0x277 // Updated for 5.4a
	EventPlay16Opcode  = 0x133 // Updated for 5.4a
	EventPlay32Opcode  = 0x17F // Updated for 5.4a
	EventPlay64Opcode  = 0x1A3 // Updated for 5.4a
	EventPlay128Opcode = 0xDF  // Updated for 5.4a
	EventPlay255Opcode = 0x2D6 // Updated for 5.4a

	MountOpcode = 0x255 // Updated for 5.4a

	WeatherChangeOpcode = 0x12D // Updated for 5.4a

	PrepareZoningOpcode = 0x2EB // Updated for 5.4a

	GaugeOpcode = 0xC2 // Updated for 5.4a

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
