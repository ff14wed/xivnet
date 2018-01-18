package datatypes

// XWorldPlayerEntry defines the player entry for a party list
type XWorldPlayerEntry struct {
	PlayerID       uint64
	U1             uint32 // Bitfield, affects the icon in various ways
	U2             uint32
	CharacterID    uint32
	ClientLanguage byte
	U3b            byte
	U3c            uint16
	Level          byte
	U4             byte
	WorldID        uint16
	ClassJob       byte
	U6             byte
	Languages      byte // Bitfield, affects available languages
	GrandCompany   byte
	Name           EntityName
}

// XWorldPartyList defines the data array for a cross world party list
type XWorldPartyList struct {
	PartyLeader uint64
	Entries     [8]XWorldPlayerEntry
}
