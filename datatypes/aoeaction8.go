package datatypes

// AoEAction8 defines the data array for an AoE action block
type AoEAction8 struct {
	ActionHeader

	EffectsList [8]ActionEffects

	U9  uint32
	U10 uint16

	Targets [8]uint64

	Position PackedPosition
	U11      uint16
	U12      uint32
}
