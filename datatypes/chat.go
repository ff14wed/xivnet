package datatypes

// Chat defines the data array for a server chat block
type Chat struct {
	ChannelID          uint64
	SpeakerCharacterID uint64
	SpeakerEntityID    uint32
	WorldID            uint16
	Flags              uint8
	SpeakerName        EntityName
	Message            ChatMessage
	Pad                uint8
}

func (Chat) IsBlockData() {}

// EgressChat defines the data array for a server chat block

type EgressChat struct {
	ChannelID uint64
	Message   ChatMessage
}

func (EgressChat) IsBlockData() {}
