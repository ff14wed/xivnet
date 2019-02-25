package datatypes_test

import (
	"encoding/json"

	"github.com/ff14wed/xivnet/datatypes"
	"github.com/ff14wed/xivnet/v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DirectorPlayScene", func() {
	var b *xivnet.Block

	BeforeEach(func() {
		b = &xivnet.Block{
			Length:    12345,
			SubjectID: 12345,
			CurrentID: 67890,
			IPCHeader: xivnet.IPCHeader{
				Opcode: datatypes.DirectorPlaySceneOpcode,
			},
			Data: xivnet.GenericBlockDataFromBytes(playSceneBlockBytes),
		}
	})

	It("parses successfully", func() {
		newB, err := datatypes.ParseBlock(b, false)
		Expect(err).ToNot(HaveOccurred())
		Expect(newB.Data).To(Equal(expectedPlaySceneBlockData))
	})

	It("marshals to JSON", func() {
		b, err := json.Marshal(expectedPlaySceneBlockData)
		Expect(err).ToNot(HaveOccurred())
		Expect(b).To(MatchJSON(expectedPlaySceneBlockDataJSON))
	})

	It("unmarshals from JSON", func() {
		d := &datatypes.DirectorPlayScene{}
		err := json.Unmarshal([]byte(expectedPlaySceneBlockDataJSON), d)
		Expect(err).ToNot(HaveOccurred())
		Expect(d).To(Equal(expectedPlaySceneBlockData))
	})

	Context("with event type 0xA0001", func() {
		var expectedCraftStateBlockData datatypes.DirectorPlayScene

		BeforeEach(func() {
			playSceneWithCraftingEvent := append(append(
				playSceneBlockBytes[:10],
				0x0A,
			), playSceneBlockBytes[11:]...)

			b = &xivnet.Block{
				Length:    12345,
				SubjectID: 12345,
				CurrentID: 67890,
				IPCHeader: xivnet.IPCHeader{
					Opcode: datatypes.DirectorPlaySceneOpcode,
				},
				Data: xivnet.GenericBlockDataFromBytes(playSceneWithCraftingEvent),
			}

			expectedCraftStateBlockData = *expectedPlaySceneBlockData
			expectedCraftStateBlockData.EventID = 0xA0001
			expectedCraftStateBlockData.Data = expectedCraftState
		})

		It("parses successfully into a DirectorPlayScene with CraftingState", func() {
			newB, err := datatypes.ParseBlock(b, false)
			Expect(err).ToNot(HaveOccurred())
			Expect(newB.Data).To(Equal(&expectedCraftStateBlockData))
		})

		It("marshals to JSON with a CraftState", func() {
			b, err := json.Marshal(expectedCraftStateBlockData)
			Expect(err).ToNot(HaveOccurred())
			Expect(b).To(MatchJSON(expectedCraftStateBlockDataJSON))
		})

		It("unmarshals from JSON with a CraftState", func() {
			d := &datatypes.DirectorPlayScene{}
			err := json.Unmarshal([]byte(expectedCraftStateBlockDataJSON), d)
			Expect(err).ToNot(HaveOccurred())
			Expect(d).To(Equal(&expectedCraftStateBlockData))
		})
	})
})
