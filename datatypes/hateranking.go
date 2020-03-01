package datatypes

// HateRankingEntry defines a single hate entry for the current target
type HateRankingEntry struct {
	ActorID uint32 // ActorID is the entity towards which the threat is directed
	HatePct byte   // HatePct is the percentage of enmity relative to the person with aggro on the current target
	Pad1    byte
	Pad2    uint16
}

// HateRanking defines the data array for the enmity list for the current target
type HateRanking struct {
	Count   byte
	U1      byte
	U2      uint16
	Entries [8]HateRankingEntry
	Pad     uint32
}

func (HateRanking) IsBlockData() {}
