package datatypes

// EventPlay2 defines the data array for a event play block
// Note: This requires an EventStart block.
// The event play header might change as we figure out more information about
// it.
type EventPlay2 struct {
	EventPlay
	U2 uint32
	U3 uint32
}

func (EventPlay2) IsBlockData() {}
