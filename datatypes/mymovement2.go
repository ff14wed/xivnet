package datatypes

// MyMovement2 defines the data array for a my movement block
type MyMovement2 struct {
	Direction float32 // 0 is South. Range [-pi,pi] <=> Counterclockwise from North
	U1        [6]uint32
	X         float32
	Z         float32
	Y         float32
}
