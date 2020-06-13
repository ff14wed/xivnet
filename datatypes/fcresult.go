package datatypes

// FreeCompanyResult defines the data array for a free company event block.
type FreeCompanyResult struct {
	FreeCompanyID     uint64
	TargetCharacterID uint64
	Type              uint32
	Result            uint32
	UpdateStatus      byte
	Identity          byte
	FreeCompanyName   FCName
	TargetName        EntityName
}

func (FreeCompanyResult) IsBlockData() {}
