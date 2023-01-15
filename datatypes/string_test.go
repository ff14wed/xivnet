package datatypes_test

import (
	"encoding/json"

	"github.com/ff14wed/xivnet/v3/datatypes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("String Type", func() {
	Describe("EntityName", func() {
		var (
			expectedName     datatypes.EntityName
			expectedUTF8Name datatypes.EntityName
		)
		BeforeEach(func() {
			expectedName = [32]byte{
				0x53, 0x74, 0x72, 0x69, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x44, 0x75, 0x6d,
				0x6d, 0x79,
			}
			expectedUTF8Name = [32]byte{
				0xe6, 0x9c, 0xa8, 0xe4, 0xba, 0xba,
			}
		})
		Describe("StringToEntityName", func() {
			It("returns the correct encoding of ASCII characters", func() {
				name := datatypes.StringToEntityName("Striking Dummy")
				Expect(name).To(Equal(expectedName))
			})
			It("returns the correct encoding of UTF-8 characters", func() {
				name := datatypes.StringToEntityName("木人")
				Expect(name).To(Equal(expectedUTF8Name))
			})
		})
		Describe("String", func() {
			BeforeEach(func() {
				expectedName[15] = 'a'
				expectedUTF8Name[7] = 'a'
			})
			It("returns the correct ASCII string", func() {
				Expect(expectedName.String()).To(Equal("Striking Dummy"))
			})
			It("returns the correct UTF-8 string", func() {
				Expect(expectedUTF8Name.String()).To(Equal("木人"))
			})
		})
		Describe("MarshalJSON", func() {
			It("marshals the string value to JSON", func() {
				jsonBytes, err := json.Marshal(expectedName)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(jsonBytes)).To(Equal(`"Striking Dummy"`))
			})
		})
		Describe("UnmarshalJSON", func() {
			It("unmarshals the string value from JSON", func() {
				var name datatypes.EntityName
				err := json.Unmarshal([]byte(`"Striking Dummy"`), &name)
				Expect(err).ToNot(HaveOccurred())
				Expect(name).To(Equal(expectedName))
			})
		})
	})

	Describe("FCTag", func() {
		var expectedTag datatypes.FCTag = [6]byte{
			0x53, 0x74, 0x72, 0x69, 0x6b,
		}
		Describe("StringToFCTag", func() {
			It("returns the correct encoding of ASCII characters", func() {
				tag := datatypes.StringToFCTag("Strik")
				Expect(tag).To(Equal(expectedTag))
			})
		})
		Describe("String", func() {
			It("returns the correct ASCII string", func() {
				Expect(expectedTag.String()).To(Equal("Strik"))
			})
		})
		Describe("MarshalJSON", func() {
			It("marshals the string value to JSON", func() {
				jsonBytes, err := json.Marshal(expectedTag)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(jsonBytes)).To(Equal(`"Strik"`))
			})
		})
		Describe("UnmarshalJSON", func() {
			It("unmarshals the string value from JSON", func() {
				var tag datatypes.FCTag
				err := json.Unmarshal([]byte(`"Strik"`), &tag)
				Expect(err).ToNot(HaveOccurred())
				Expect(tag).To(Equal(expectedTag))
			})
		})
	})

	Describe("FCName", func() {
		var expectedName datatypes.FCName = [46]byte{
			0x53, 0x74, 0x72, 0x69, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x44, 0x75, 0x6d,
			0x6d, 0x79,
		}
		Describe("StringToFCName", func() {
			It("returns the correct encoding of ASCII characters", func() {
				name := datatypes.StringToFCName("Striking Dummy")
				Expect(name).To(Equal(expectedName))
			})
		})
		Describe("String", func() {
			It("returns the correct ASCII string", func() {
				Expect(expectedName.String()).To(Equal("Striking Dummy"))
			})
		})
		Describe("MarshalJSON", func() {
			It("marshals the string value to JSON", func() {
				jsonBytes, err := json.Marshal(expectedName)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(jsonBytes)).To(Equal(`"Striking Dummy"`))
			})
		})
		Describe("UnmarshalJSON", func() {
			It("unmarshals the string value from JSON", func() {
				var name datatypes.FCName
				err := json.Unmarshal([]byte(`"Striking Dummy"`), &name)
				Expect(err).ToNot(HaveOccurred())
				Expect(name).To(Equal(expectedName))
			})
		})
	})

	Describe("ChatMessage", func() {
		var (
			expectedChatMessage datatypes.ChatMessage
			expectedUTF8Message datatypes.ChatMessage
		)
		BeforeEach(func() {
			expectedChatMessage = [1024]byte{
				0x53, 0x74, 0x72, 0x69, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x44, 0x75, 0x6d,
				0x6d, 0x79,
			}
			expectedUTF8Message = [1024]byte{
				0xe6, 0x9c, 0xa8, 0xe4, 0xba, 0xba,
			}
		})
		Describe("StringToChatMessage", func() {
			It("returns the correct encoding of ASCII characters", func() {
				msg := datatypes.StringToChatMessage("Striking Dummy")
				Expect(msg).To(Equal(expectedChatMessage))
			})
			It("returns the correct encoding of UTF-8 characters", func() {
				msg := datatypes.StringToChatMessage("木人")
				Expect(msg).To(Equal(expectedUTF8Message))
			})
		})
		Describe("String", func() {
			BeforeEach(func() {
				expectedChatMessage[15] = 'a'
				expectedUTF8Message[7] = 'a'
			})
			It("returns the correct ASCII string", func() {
				Expect(expectedChatMessage.String()).To(Equal("Striking Dummy"))
			})
			It("returns the correct UTF-8 string", func() {
				Expect(expectedUTF8Message.String()).To(Equal("木人"))
			})
		})
		Describe("MarshalJSON", func() {
			It("marshals the string value to JSON", func() {
				jsonBytes, err := json.Marshal(expectedChatMessage)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(jsonBytes)).To(Equal(`"Striking Dummy"`))
			})
		})
		Describe("UnmarshalJSON", func() {
			It("unmarshals the string value from JSON", func() {
				var message datatypes.ChatMessage
				err := json.Unmarshal([]byte(`"Striking Dummy"`), &message)
				Expect(err).ToNot(HaveOccurred())
				Expect(message).To(Equal(expectedChatMessage))
			})
		})
	})
})
