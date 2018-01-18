package datatypes

// Target defines the data array for a target block
type Target struct {
	U1       uint32
	U2       uint64
	U3       uint64
	U4       uint32
	TargetID uint32 // 0xE0000000 means deselect target
	U5       uint32
}
