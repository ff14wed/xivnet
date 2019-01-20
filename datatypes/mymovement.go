package datatypes

// MyMovement defines the data array for a my movement block
type MyMovement struct {
	Direction float32 // 0 is South. Range [-pi,pi] <=> Counterclockwise from North
	U1        uint32
	U2        uint32
	X         float32
	Z         float32
	Y         float32
}

func (MyMovement) IsBlockData() {}
