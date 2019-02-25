package datatypes

// DirectorPlayScene defines the data array for a event scene block
type DirectorPlayScene struct {
	CharID            uint64
	U1                [9]uint32
	CraftAction       uint32
	U2                uint32
	StepNum           uint32
	TotalProgress     uint32
	ProgressDelta     int32
	TotalQuality      uint32
	QualityDelta      int32
	HQChance          uint32
	Durability        uint32
	DurabilityDelta   int32
	CurrentCondition  uint32
	PreviousCondition uint32
	U6                [17]uint32
}

func (DirectorPlayScene) IsBlockData() {}
