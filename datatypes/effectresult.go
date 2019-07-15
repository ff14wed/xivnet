package datatypes

// EffectResultEntry is an entry of EffectResult
type EffectResultEntry struct {
	Index    byte // Position of status effect
	U6       byte
	EffectID uint16
	Param    uint16
	U7       uint16
	Duration float32
	ActorID  uint32
}

// EffectResult defines the data array for an effect result.
type EffectResult struct {
	U1      uint32
	ActorID uint32

	CurrentHP uint32
	MaxHP     uint32
	CurrentMP uint16
	CurrentTP uint16
	MaxMP     uint16

	Pad2     byte
	ClassJob byte

	Pad3  byte
	Count byte

	U4      uint16
	Entries [4]EffectResultEntry
	U8      uint32
}

func (EffectResult) IsBlockData() {}
