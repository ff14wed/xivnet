package datatypes

// Notify144 defines the data array for a notify4 block
type Notify144 struct {
	Type     uint16
	Pad1     uint16
	P1       uint32
	P2       uint32
	P3       uint32
	P4       uint32
	Pad2     uint32
	TargetID uint32 // 0xE0000000 means deselect target
	U1       uint32
}

func (Notify144) IsBlockData() {}
