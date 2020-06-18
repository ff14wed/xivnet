package datatypes_test

import (
	"encoding/json"

	"github.com/ff14wed/xivnet/v3"
	"github.com/ff14wed/xivnet/v3/datatypes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EventPlay32", func() {
	var b *xivnet.Block

	BeforeEach(func() {
		b = &xivnet.Block{
			Length:    12345,
			SubjectID: 12345,
			CurrentID: 67890,
			IPCHeader: xivnet.IPCHeader{
				ServerID: 123,
				Opcode:   datatypes.EventPlay32Opcode,
			},
			Data: xivnet.GenericBlockDataFromBytes(eventPlay32BlockBytes),
		}
	})

	It("parses successfully", func() {
		newB, err := datatypes.ParseBlock(b, false)
		Expect(err).ToNot(HaveOccurred())
		Expect(newB.Data).To(Equal(expectedEventPlay32BlockData))
	})

	It("marshals to JSON", func() {
		jsonBytes, err := json.Marshal(expectedEventPlay32BlockData)
		Expect(err).ToNot(HaveOccurred())
		Expect(jsonBytes).To(MatchJSON(expectedEventPlay32BlockDataJSON))
	})

	It("unmarshals from JSON", func() {
		d := &datatypes.EventPlay32{}
		err := json.Unmarshal([]byte(expectedEventPlay32BlockDataJSON), d)
		Expect(err).ToNot(HaveOccurred())
		Expect(d).To(Equal(expectedEventPlay32BlockData))
	})

	Context("with event type 0xA0001", func() {
		var expectedCraftStateBlockData datatypes.EventPlay32

		BeforeEach(func() {
			eventPlay32WithCraftingEvent := make([]byte, len(eventPlay32BlockBytes))
			copy(eventPlay32WithCraftingEvent, eventPlay32BlockBytes[:10])
			eventPlay32WithCraftingEvent[10] = 0x0A
			copy(eventPlay32WithCraftingEvent[11:], eventPlay32BlockBytes[11:])

			b = &xivnet.Block{
				Length:    12345,
				SubjectID: 12345,
				CurrentID: 67890,
				IPCHeader: xivnet.IPCHeader{
					ServerID: 123,
					Opcode:   datatypes.EventPlay32Opcode,
				},
				Data: xivnet.GenericBlockDataFromBytes(eventPlay32WithCraftingEvent),
			}

			expectedCraftStateBlockData = *expectedEventPlay32BlockData
			expectedCraftStateBlockData.EventID = 0xA0001
			expectedCraftStateBlockData.Data = expectedCraftState
		})

		It("parses successfully into a EventPlay32 with CraftingState", func() {
			newB, err := datatypes.ParseBlock(b, false)
			Expect(err).ToNot(HaveOccurred())
			Expect(newB.Data).To(Equal(&expectedCraftStateBlockData))
		})

		It("marshals to JSON with a CraftState", func() {
			jsonBytes, err := json.Marshal(expectedCraftStateBlockData)
			Expect(err).ToNot(HaveOccurred())
			Expect(jsonBytes).To(MatchJSON(expectedCraftStateBlockDataJSON))
		})

		It("unmarshals from JSON with a CraftState", func() {
			d := &datatypes.EventPlay32{}
			err := json.Unmarshal([]byte(expectedCraftStateBlockDataJSON), d)
			Expect(err).ToNot(HaveOccurred())
			Expect(d).To(Equal(&expectedCraftStateBlockData))
		})
	})
})
