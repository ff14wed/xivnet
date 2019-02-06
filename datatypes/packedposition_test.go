package datatypes_test

import (
	"encoding/json"

	"github.com/ff14wed/xivnet/datatypes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Packed Position Type", func() {
	Describe("PackedCoord", func() {
		Describe("Float", func() {
			It("returns the correct number", func() {
				pc := datatypes.PackedCoord(0x259B)
				Expect(pc.Float()).To(BeNumerically("~", -706.19354, 0.01))
			})
		})
		Describe("SetFloat", func() {
			It("sets the correct number", func() {
				var pc datatypes.PackedCoord
				pc.SetFloat(-706.19354)
				Expect(pc).To(Equal(datatypes.PackedCoord(0x259B)))
			})
		})
		Describe("MarshalJSON", func() {
			It("marshals the floating point values to JSON", func() {
				var pc datatypes.PackedCoord
				pc.SetFloat(-706.19354)
				jsonBytes, err := json.Marshal(pc)
				Expect(err).ToNot(HaveOccurred())
				Expect(string(jsonBytes)).To(ContainSubstring("-706"))
			})
		})
		Describe("UnmarshalJSON", func() {
			It("unmarshals the floating point value from JSON", func() {
				var pc datatypes.PackedCoord
				err := json.Unmarshal([]byte("-706.19354"), &pc)
				Expect(err).ToNot(HaveOccurred())
				Expect(pc).To(Equal(datatypes.PackedCoord(0x259B)))
			})
		})
	})
	Describe("PackedPosition", func() {
		It("returns the correct coordinates", func() {
			pp := datatypes.PackedPosition{
				X: 0x259B,
				Y: 0x83E7,
				Z: 0x204E,
			}
			Expect(pp.X.Float()).To(BeNumerically("~", -706.19354, 0.01))
			Expect(pp.Y.Float()).To(BeNumerically("~", 30.5, 0.01))
			Expect(pp.Z.Float()).To(BeNumerically("~", -747.62988, 0.02))
		})
		It("sets the correct coordinates", func() {
			pp := datatypes.PackedPosition{}
			pp.X.SetFloat(-706.19354)
			pp.Y.SetFloat(30.5)
			pp.Z.SetFloat(-747.62988)
			Expect(pp).To(Equal(datatypes.PackedPosition{
				X: 0x259B,
				Y: 0x83E7,
				Z: 0x204E,
			}))
		})
		It("marshals to JSON", func() {
			pp := datatypes.PackedPosition{
				X: 0x259B,
				Y: 0x83E7,
				Z: 0x204E,
			}
			jsonBytes, err := json.Marshal(pp)
			Expect(err).ToNot(HaveOccurred())
			Expect(string(jsonBytes)).To(ContainSubstring("-706"))
			Expect(string(jsonBytes)).To(ContainSubstring("30.5"))
			Expect(string(jsonBytes)).To(ContainSubstring("-747.6"))
		})

		It("unmarshals from JSON", func() {
			var pp datatypes.PackedPosition
			err := json.Unmarshal(
				[]byte(`{"X":-706.19354,"Y":30.5,"Z":-747.62988}`),
				&pp,
			)
			Expect(err).ToNot(HaveOccurred())
			Expect(pp).To(Equal(datatypes.PackedPosition{
				X: 0x259B,
				Y: 0x83E7,
				Z: 0x204E,
			}))
		})
	})
})
