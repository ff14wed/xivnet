package datatypes

// Chat defines the data array for a server chat block for cross-world chat
// Seems to only be used for cross-world linkshells
type ChatXWorld struct {
	ChannelID          uint64
	SpeakerCharacterID uint64
	SpeakerEntityID    uint32
	WorldID            uint16
	WorldID2           uint16
	Flags              uint8
	SpeakerName        EntityName
	Message            ChatMessage

	Pad [7]uint8
}

func (ChatXWorld) IsBlockData() {}

// EgressChatXWorld defines the data array for a server chat block for cross-world
// chat
type EgressChatXWorld struct {
	ChannelID uint64
	Flags     uint8
	Message   ChatMessage

	Pad [7]uint8
}

func (EgressChatXWorld) IsBlockData() {}
