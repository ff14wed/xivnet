package datatypes

// ChatZone defines the data array for a zone chat block
type ChatZone struct {
	CharacterID uint64
	EntityID    uint32
	WorldID     uint16
	Type        uint16
	SpeakerName EntityName
	Message     ChatMessage
}

func (ChatZone) IsBlockData() {}

// EgressChatZone defines the data array for a zone chat block

type EgressChatZone struct {
	ClientTime uint32
	Position   EgressChatPosition
	Type       uint16
	Message    ChatMessage
	Pad1       uint16
	Pad2       uint32
}

func (EgressChatZone) IsBlockData() {}

type EgressChatPosition struct {
	EntityID  uint32
	X, Y, Z   float32
	Direction float32
}
