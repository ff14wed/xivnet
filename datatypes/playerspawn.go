package datatypes

// ModelInfo enumerates all of the different options for customizing
// the character model
type ModelInfo struct {
	Race                byte // 2 Elezen, 3 Lalafell, 4 Miqo'te, 5 Roe, 6 Au Ra, else Hyur
	Gender              byte // 0 is male, 1 is female
	BodyType            byte // CHANGE AT OWN RISK
	Height              byte // Scale from 0-100
	Tribe               byte
	Face                byte
	Hairstyle           byte
	HairHighlight       byte // 8th bit toggles highlight
	SkinTone            byte
	OddEyeColor         byte
	HairColor           byte
	HairHighlightColor  byte
	FacialFeatures      byte // Race specific toggles. i.e. 6th bit toggles right limbal ring. 7th bit toggles left limbal ring.
	FacialFeaturesColor byte
	Eyebrows            byte
	EyeColor            byte
	EyeShape            byte
	Nose                byte
	Jaw                 byte
	Mouth               byte // Bitfield toggles light/dark/none lip color
	LipColor            byte
	TailLength          byte // Scale from 1-100
	TailType            byte
	BustSize            byte // Scale from 1-100
	FacePaintType       byte
	FacePaintColor      byte
}

// MountInfo enumerates the fields for mount information
type MountInfo struct {
	ID    byte
	Head  byte
	Body  byte
	Feet  byte
	Color byte
}

// PlayerSpawn defines the data array for a new entity block
type PlayerSpawn struct {
	Title                   uint16
	U1b                     uint16
	U2b                     byte
	U2ab                    byte
	GMRank                  byte
	U3b                     byte
	MobAggressive           byte
	OnlineStatus, Pose, U3d byte
	U4                      uint32
	TargetID                uint64

	U6, U7    uint32
	Weapon    WeaponGear
	CraftTool WeaponGear

	BNPCBase, BNPCName uint32
	U18, U19, U20      uint32
	OwnerID            uint32
	UnkID3             uint32

	CurrentHP, MaxHP uint32
	DisplayFlags     uint32
	FateID           uint16
	CurrentMP        uint16
	CurrentTP        uint16
	MaxMP            uint16

	U21a       uint16
	ModelChara uint16
	Direction  uint16 // Quantized direction 0x0000 ~ 0xFFFF, NWSE <=> 0,0x4000,0x8000,0xC000
	Minion     uint16
	Index      byte
	State      byte // 0-1 for alive, 2 for dead, 3 for persistent emote
	Emote      byte // Applies for when State is 3
	Type       byte // 1 for player, 2 for NPC, else furniture
	Subtype    byte // 4 for players, 2 pet, 3 companion, 5 mob, 7 minion
	U25b       byte
	U25c       uint16

	EnemyType byte // 0 for friendly, anything else is an enemy
	Level     byte
	ClassJob  byte
	U26d      byte
	U27a      uint16

	MountInfo MountInfo
	U28b      byte
	U28c      uint32
	U29       uint32
	Statuses  [30]StatusEffect
	X         float32
	Z         float32
	Y         float32
	Head      Gear
	Body      Gear
	Hand      Gear
	Leg       Gear
	Foot      Gear
	Ear       Gear
	Neck      Gear
	Wrist     Gear
	Ring1     Gear
	Ring2     Gear

	Name EntityName

	Model ModelInfo

	FCTag FCTag
	U30   uint64
}
