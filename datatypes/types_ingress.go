package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x26B // Updated for 5.45
	InitZoneOpcode      = 0x1C6 // Updated for 5.45
	ControlOpcode       = 0xF0  // Updated for 5.45
	ControlSelfOpcode   = 0x350 // Updated for 5.45
	ControlTargetOpcode = 0x29A // Updated for 5.45
	RemoveEntityOpcode  = 0x377 // Updated for 5.45
	UpdateHPMPTPOpcode  = 0x1DB // Updated for 5.45

	ChatZoneOpcode = 0x7B // Updated for 5.45

	UpdateStatusesOpcode       = 0xC2 // Updated for 5.45
	UpdateStatusesEurekaOpcode = 0x8B // Updated for 5.45
	UpdateStatusesBossOpcode   = 0x91 // Updated for 5.45

	ActionOpcode      = 0x21F // Updated for 5.45
	AoEAction8Opcode  = 0x3DF // Updated for 5.45
	AoEAction16Opcode = 0xAD  // Updated for 5.45
	AoEAction24Opcode = 0x229 // Updated for 5.45
	AoEAction32Opcode = 0x197 // Updated for 5.45

	ObjectSpawnOpcode = 0x3C8 // Updated for 5.45
	PlayerSpawnOpcode = 0x83  // Updated for 5.45
	NPCSpawnOpcode    = 0x3C4 // Updated for 5.45
	NPCSpawn2Opcode   = 0x11C // Updated for 5.45

	MovementOpcode = 0x3A9 // Updated for 5.45
	SetPosOpcode   = 0x283 // Updated for 5.45

	CastingOpcode = 0x2B2 // Updated for 5.45

	HateRankingOpcode = 0x144 // Updated for 5.45
	HateListOpcode    = 0xF8  // Updated for 5.45

	EquipChangeOpcode = 0x375 // Updated for 5.45

	EventPlayOpcode    = 0x1D2 // Updated for 5.45
	EventPlay4Opcode   = 0x2CF // Updated for 5.45
	EventPlay8Opcode   = 0x3B2 // Updated for 5.45
	EventPlay16Opcode  = 0x1AE // Updated for 5.45
	EventPlay32Opcode  = 0x11D // Updated for 5.45
	EventPlay64Opcode  = 0x262 // Updated for 5.45
	EventPlay128Opcode = 0x318 // Updated for 5.45
	EventPlay255Opcode = 0x37F // Updated for 5.45

	MountOpcode = 0x1F8 // Updated for 5.45

	WeatherChangeOpcode = 0x250 // Updated for 5.45

	PrepareZoningOpcode = 0xAB // Updated for 5.45

	GaugeOpcode = 0x278 // Updated for 5.45

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
