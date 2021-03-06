package datatypes

// EgressInstanceMovement defines the data array for a outgoing movement block inside an instance
type EgressInstanceMovement struct {
	Direction float32 // 0 is South. Range [-pi,pi] <=> Counterclockwise from North
	U1        [5]uint32
	X         float32
	Y         float32
	Z         float32
	U2        uint32
}

func (EgressInstanceMovement) IsBlockData() {}
