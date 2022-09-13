package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x12f // Updated for 6.21
	InitZoneOpcode      = 0x36A // Updated for 6.21
	ControlOpcode       = 0x335 // Updated for 6.21
	ControlSelfOpcode   = 0x2ED // Updated for 6.21
	ControlTargetOpcode = 0x212 // Updated for 6.21
	RemoveEntityOpcode  = 0x223 // Updated for 6.21
	UpdateHPMPTPOpcode  = 0x193 // Updated for 6.21

	ChatZoneOpcode = 0x2EC // Updated for 6.21

	UpdateStatusesOpcode       = 0xFA  // Updated for 6.21
	UpdateStatusesEurekaOpcode = 0x160 // Updated for 6.21
	UpdateStatusesBossOpcode   = 0x169 // Updated for 6.21

	ActionOpcode      = 0x1CD // Updated for 6.21
	AoEAction8Opcode  = 0x35C // Updated for 6.21
	AoEAction16Opcode = 0xB0  // Updated for 6.21
	AoEAction24Opcode = 0x276 // Updated for 6.21
	AoEAction32Opcode = 0x1F5 // Updated for 6.21

	ObjectSpawnOpcode = 0x281 // Updated for 6.21
	PlayerSpawnOpcode = 0x29D // Updated for 6.21
	NPCSpawnOpcode    = 0x2E4 // Updated for 6.21
	NPCSpawn2Opcode   = 0x3AA // Updated for 6.21

	MovementOpcode = 0x39F // Updated for 6.21
	SetPosOpcode   = 0xC6  // Updated for 6.21

	CastingOpcode = 0x2E8 // Updated for 6.21

	HateRankingOpcode = 0x225 // Updated for 6.21
	HateListOpcode    = 0x39B // Updated for 6.21

	PlayerStatsOpcode = 0x347 // Updated for 6.21

	EquipChangeOpcode = 0x17B // Updated for 6.21

	EventPlayOpcode    = 0x2DB // Updated for 6.21
	EventPlay4Opcode   = 0x103 // Updated for 6.21
	EventPlay8Opcode   = 0x68  // Updated for 6.21
	EventPlay16Opcode  = 0x106 // Updated for 6.21
	EventPlay32Opcode  = 0x2F9 // Updated for 6.21
	EventPlay64Opcode  = 0xC4  // Updated for 6.21
	EventPlay128Opcode = 0x218 // Updated for 6.21
	EventPlay255Opcode = 0xEC  // Updated for 6.21

	MountOpcode = 0x2CF // Updated for 6.21

	WeatherChangeOpcode = 0x162 // Updated for 6.21

	PrepareZoningOpcode = 0x121 // Updated for 6.21

	GaugeOpcode = 0x21C // Updated for 6.21

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

	registerInBlockFactory(PlayerStatsOpcode, func() xivnet.BlockData { return new(PlayerStats) })

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
