package datatypes

// AoEAction24 defines the data array for an AoE action block
type AoEAction24 struct {
	ActionHeader

	EffectsList [24]ActionEffects

	U9  uint32
	U10 uint16

	Targets [24]uint64

	Position PackedPosition
	U11      uint16
	U12      uint32
}
