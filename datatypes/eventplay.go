package datatypes

type EventPlayHeader struct {
	ActorID    uint64
	EventID    uint32 // 0xA0001 for crafting
	Scene      uint16
	Pad1       uint16
	Flags      uint32
	P1         uint32
	ParamCount byte
	Pad2       [3]byte
	P2         uint32
}

// EventPlay defines the data array for a event play block
// Note: This requires an EventStart block.
type EventPlay struct {
	EventPlayHeader
	Params [2]uint32
}

func (EventPlay) IsBlockData() {}

// EventPlay4 defines the data array for a event play block with 4 params
// Note: This requires an EventStart block.
type EventPlay4 struct {
	EventPlayHeader
	Params [4]uint32
}

func (EventPlay4) IsBlockData() {}
