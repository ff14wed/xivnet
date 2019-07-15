package datatypes

// EventPlay defines the data array for a event play block
// Note: This requires an EventStart block.
// The event play header might change as we figure out more information about
// it.
type EventPlay struct {
	ActorID uint64
	EventID uint32 // 0xA0001 for crafting
	Scene   uint16
	Pad1    uint16
	Flags   uint32
	P1      uint32
	P2      byte
	Pad2    [3]byte
	P3      uint32
	P4      uint32
	P5      uint32
}

func (EventPlay) IsBlockData() {}
