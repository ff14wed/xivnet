package datatypes_test

import (
	"encoding/json"

	"github.com/ff14wed/xivnet/v3"
	"github.com/ff14wed/xivnet/v3/datatypes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("EventPlay", func() {
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

	It("random data fails to unmarshal to a craft state", func() {
		_, err := datatypes.UnmarshalCraftState(expectedEventPlay32BlockData)
		Expect(err).To(MatchError(ContainSubstring("not a craft state")))
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
		})

		It("parses successfully into a EventPlay32", func() {
			newB, err := datatypes.ParseBlock(b, false)
			Expect(err).ToNot(HaveOccurred())
			Expect(newB.Data).To(Equal(&expectedCraftStateBlockData))
		})

		It("marshals the data to a CraftState", func() {
			cs, err := datatypes.UnmarshalCraftState(&expectedCraftStateBlockData)
			Expect(err).ToNot(HaveOccurred())
			Expect(cs).To(Equal(&expectedCraftState))

			data, err := datatypes.MarshalCraftState(cs)
			Expect(err).ToNot(HaveOccurred())
			Expect(data).To(Equal(expectedCraftStateBlockData.Data[:]))
		})

		It("marshals into a CraftState from an EventPlay64", func() {
			eventPlay64 := &datatypes.EventPlay64{
				EventPlayHeader: expectedEventPlay32BlockData.EventPlayHeader,
			}
			copy(eventPlay64.Data[:], expectedEventPlay32BlockData.Data[:])
			cs, err := datatypes.UnmarshalCraftState(&expectedCraftStateBlockData)
			Expect(err).ToNot(HaveOccurred())
			Expect(cs.Progress).To(Equal(uint32(256)))
		})
	})
})
