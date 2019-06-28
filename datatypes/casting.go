package datatypes

// Casting defines the data array for a casting block
type Casting struct {
	ActionIDName uint16
	U1           uint16
	ActionID     uint32
	CastTime     float32
	TargetID     uint32
	Direction    uint16 // Quantized direction 0x0000 ~ 0xFFFF, NWSE <=> 0,0x4000,0x8000,0xC000
	U2           uint16
	UnkID1       uint32
	Position     PackedPosition
	U3           uint16
}

func (Casting) IsBlockData() {}
