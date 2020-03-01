package datatypes

// UpdateHPMPTP defines the data array for a updateHPMPTP block
type UpdateHPMPTP struct {
	HP  uint32
	MP  uint16
	Pad uint16
}

func (UpdateHPMPTP) IsBlockData() {}
