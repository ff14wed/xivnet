package datatypes

// UpdateHPMPTP defines the data array for a updateHPMPTP block
type UpdateHPMPTP struct {
	HP uint32
	MP uint16
	TP uint16
	U1 uint32
	U2 uint32
}

func (UpdateHPMPTP) IsBlockData() {}
