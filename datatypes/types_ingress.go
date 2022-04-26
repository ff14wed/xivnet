package datatypes

import "github.com/ff14wed/xivnet/v3"

var inTypeRegistry = make(map[uint16]func() xivnet.BlockData)

// Opcodes that define the datatypes of incoming (from server) network blocks
const (
	EffectResultOpcode  = 0x309 // Updated for 6.11
	InitZoneOpcode      = 0x1F5 // Updated for 6.11
	ControlOpcode       = 0x16F // Updated for 6.11
	ControlSelfOpcode   = 0x3AE // Updated for 6.11
	ControlTargetOpcode = 0x1B2 // Updated for 6.11
	RemoveEntityOpcode  = 0xAD  // Updated for 6.11
	UpdateHPMPTPOpcode  = 0x28B // Updated for 6.11

	ChatZoneOpcode = 0x1C7 // Updated for 6.11

	UpdateStatusesOpcode       = 0xD2  // Updated for 6.11
	UpdateStatusesEurekaOpcode = 0x182 // Updated for 6.11
	UpdateStatusesBossOpcode   = 0x2B4 // Updated for 6.11

	ActionOpcode      = 0x398 // Updated for 6.11
	AoEAction8Opcode  = 0x359 // Updated for 6.11
	AoEAction16Opcode = 0x260 // Updated for 6.11
	AoEAction24Opcode = 0x209 // Updated for 6.11
	AoEAction32Opcode = 0x39F // Updated for 6.11

	ObjectSpawnOpcode = 0x2D3 // Updated for 6.11
	PlayerSpawnOpcode = 0x18F // Updated for 6.11
	NPCSpawnOpcode    = 0x2B1 // Updated for 6.11
	NPCSpawn2Opcode   = 0x83  // Updated for 6.11

	MovementOpcode = 0x397 // Updated for 6.11
	SetPosOpcode   = 0x317 // Updated for 6.11

	CastingOpcode = 0x163 // Updated for 6.11

	HateRankingOpcode = 0x1C5 // Updated for 6.11
	HateListOpcode    = 0x24C // Updated for 6.11

	EquipChangeOpcode = 0x17A // Updated for 6.11

	EventPlayOpcode    = 0x313 // Updated for 6.11
	EventPlay4Opcode   = 0x36D // Updated for 6.11
	EventPlay8Opcode   = 0x364 // Updated for 6.11
	EventPlay16Opcode  = 0x9E  // Updated for 6.11
	EventPlay32Opcode  = 0x3C8 // Updated for 6.11
	EventPlay64Opcode  = 0x16B // Updated for 6.11
	EventPlay128Opcode = 0x2FB // Updated for 6.11
	EventPlay255Opcode = 0xD8  // Updated for 6.11

	MountOpcode = 0x268 // Updated for 6.11

	WeatherChangeOpcode = 0x2C6 // Updated for 6.11

	PrepareZoningOpcode = 0x1CF // Updated for 6.11

	GaugeOpcode = 0x2A7 // Updated for 6.11

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
