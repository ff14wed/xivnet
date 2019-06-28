package datatypes

// EgressMovement defines the data array for an outgoing movement block
type EgressMovement struct {
	Direction float32 // 0 is South. Range [-pi,pi] <=> Counterclockwise from North
	U1        uint32
	X         float32
	Y         float32
	Z         float32
	U2        uint32
}

func (EgressMovement) IsBlockData() {}
