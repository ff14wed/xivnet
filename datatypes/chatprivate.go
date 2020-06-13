package datatypes

// ChatFrom defines the data array for a private chat block
type ChatFrom struct {
	FromCharacterID uint64
	WorldID         uint16
	Flags           uint8
	FromName        EntityName
	Message         ChatMessage

	Pad  uint8
	Pad2 uint32
}

func (ChatFrom) IsBlockData() {}

// ChatTo defines the data array for a private chat block
type ChatTo struct {
	ToCharacterID uint64
	ChannelID     uint64
	WorldID       uint16
	Flags         uint8
	ToName        EntityName
	Message       ChatMessage

	Pad  uint8
	Pad2 uint32
}

func (ChatTo) IsBlockData() {}
