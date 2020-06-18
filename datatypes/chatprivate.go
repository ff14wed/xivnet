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

// ChatFromXWorld defines the data array for a private chat block cross world
type ChatFromXWorld struct {
	FromCharacterID uint64
	WorldID         uint16
	U1              uint16
	U2              uint32
	FromEntityID    uint32
	WorldID2        uint16
	U3              uint16
	Flags           uint8
	FromName        EntityName
	Message         ChatMessage

	Pad  uint8
	Pad2 uint32
	Pad3 uint16
}

func (ChatFromXWorld) IsBlockData() {}

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
