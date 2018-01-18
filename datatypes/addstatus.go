package datatypes

// AddStatusEntry is an entry of AddStatus
type AddStatusEntry struct {
	Index    byte // Position of status effect
	U6       byte
	EffectID uint16
	Extra    uint16
	U7       uint16
	Duration float32
	ActorID  uint32
}

// AddStatus defines the data array for an add status effect block
// This block is used more for updating HP, MP, TP than adding a status effect
type AddStatus struct {
	U1        uint32
	ActorID   uint32
	U2, U3    byte
	Pad1      uint16
	CurrentHP uint32
	CurrentMP uint16
	CurrentTP uint16
	MaxHP     uint32
	MaxMP     uint16
	Count     byte
	U5        byte
	Entries   [4]AddStatusEntry
	U8        uint32
}
