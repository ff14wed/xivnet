package datatypes

// HateRankingEntry defines a single hate entry for the current target
type HateRankingEntry struct {
	ActorID uint32 // ActorID is the entity towards which the threat is directed
	Hate    uint32 // Hate is the total amount of enmity accrued by ActorID
}

// HateRanking defines the data array for the enmity list for the current target
type HateRanking struct {
	Count   byte
	U1      byte
	U2      uint16
	Entries [32]HateRankingEntry
	Pad     uint32
}

func (HateRanking) IsBlockData() {}
