package datatypes

// MyAction defines the data array for a my action block
type MyAction struct {
	Type uint16 // 3 means P1 is target
	U1   uint16
	P1   uint32
	P2   uint32
	P3   uint32
	P4   uint32
	P5   uint32
	P6   uint32
	P7   uint32
}

func (MyAction) IsBlockData() {}
