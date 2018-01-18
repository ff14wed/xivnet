package datatypes

// Casting defines the data array for a casting block
type Casting struct {
	ActionIDName uint16
	U1           uint16
	ActionID     uint32
	CastTime     float32
	TargetID     uint32
	Direction    float32
	UnkID1       uint32
	Position     PackedPosition
	U3           uint16
}
