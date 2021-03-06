package datatypes

// StatusEffect represents the data structure for a status effect
type StatusEffect struct {
	ID       uint16
	Param    uint16
	Duration float32
	ActorID  uint32
}
