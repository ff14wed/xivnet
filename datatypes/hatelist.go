package datatypes

// HateEntry defines the a full hate list entry
type HateEntry struct {
	EnemyID uint32 // OwnerID is the entity that has enmity against the Actor
	HatePct byte   // HatePct is how close to aggro aggro this mob (100% is aggro)
	U1      byte   // U1 and U2 are leftovers from reuse of the HateRanking struct
	U2      uint16
}

// HateList defines the data array for a hate list
type HateList struct {
	Count   byte
	U1      byte
	U2      uint16
	Entries [32]HateEntry
	Pad     uint32
}
