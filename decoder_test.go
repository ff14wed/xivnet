package xivnet_test

import (
	"bytes"

	"github.com/ff14wed/xivnet/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"github.com/onsi/gomega/types"
)

func matchExpectedFrame(expectedFrame xivnet.Frame) types.GomegaMatcher {
	return gstruct.PointTo(gstruct.MatchFields(gstruct.IgnoreExtras, gstruct.Fields{
		"Preamble":           Equal(expectedFrame.Preamble),
		"Time":               Equal(expectedFrame.Time),
		"Length":             Equal(expectedFrame.Length),
		"ConnectionType":     Equal(expectedFrame.ConnectionType),
		"Count":              Equal(expectedFrame.Count),
		"Reserved1":          Equal(expectedFrame.Reserved1),
		"Compression":        Equal(expectedFrame.Compression),
		"Reserved2":          Equal(expectedFrame.Reserved2),
		"DecompressedLength": Equal(expectedFrame.DecompressedLength),
		"Blocks":             Equal(expectedFrame.Blocks),
	}))
}

var _ = Describe("Decoder", func() {
	Describe("Decode", func() {
		Context("with zlib compressed blocks", func() {
			It("properly decodes a packet into the correct structures", func() {
				buf := bytes.NewReader(zlibPacket)
				d := xivnet.NewDecoder(buf, 32768)
				f, err := d.NextFrame()
				Expect(err).ToNot(HaveOccurred())
				Expect(f).To(matchExpectedFrame(expectedZlibFrame))
			})
		})

		Context("with a zero-block zlib compressed frame", func() {
			It("correct returns a frame with no blocks", func() {
				byteBuf := bytes.NewBuffer(zeroBlockPacket)
				d := xivnet.NewDecoder(byteBuf, 32768)
				f, err := d.NextFrame()
				Expect(err).ToNot(HaveOccurred())
				Expect(f.Count).To(BeZero())
				Expect(f.Blocks).To(BeEmpty())
			})
		})

		Context("with multiple frames on the buffer", func() {
			It("decodes the packets until there is nothing left on the buffer", func() {
				buf := bytes.NewBuffer(append(zlibPacket, zlibPacket...))

				d := xivnet.NewDecoder(buf, 32768)
				for i := 0; i < 2; i++ {
					f, err := d.NextFrame()
					Expect(err).ToNot(HaveOccurred())
					Expect(f).To(matchExpectedFrame(expectedZlibFrame))
				}

				_, err := d.NextFrame()
				Expect(err).To(HaveOccurred())
			})
		})

		Context("with a non-zlib compressed packet and short block data", func() {
			It("properly decodes a packet into the correct structures", func() {
				buf := bytes.NewBuffer(nonZlibPacket)
				d := xivnet.NewDecoder(buf, 32768)
				frame, err := d.NextFrame()
				Expect(err).ToNot(HaveOccurred())
				Expect(frame).To(matchExpectedFrame(expectedNonZlibFrame))
			})
		})

		Context("with incomplete data on the buffer", func() {
			It("eventually is able to read the frame", func() {
				byteBuf := bytes.NewBuffer(nil)
				_, err := byteBuf.Write(zlibPacket[:69])
				Expect(err).ToNot(HaveOccurred())

				d := xivnet.NewDecoder(byteBuf, 32768)
				_, err = d.NextFrame()
				Expect(err).To(MatchError("peeking data failed reading 148 bytes from buffer: EOF"))

				_, err = byteBuf.Write(zlibPacket[69:])
				Expect(err).ToNot(HaveOccurred())

				f, err := d.NextFrame()
				Expect(err).ToNot(HaveOccurred())

				Expect(f).To(matchExpectedFrame(expectedZlibFrame))
			})
		})

		Context("with a decoder that has too small a buffer", func() {
			It("returns an error", func() {
				buf := bytes.NewReader(zlibPacket)
				d := xivnet.NewDecoder(buf, 16)
				_, err := d.NextFrame()
				Expect(err).To(MatchError("invalid frame length: 28 (max 16)"))
			})
			Context("with a slightly larger but still too small of a buffer", func() {
				It("returns an error", func() {
					buf := bytes.NewReader(zlibPacket)
					d := xivnet.NewDecoder(buf, 40)
					_, err := d.NextFrame()
					Expect(err).To(MatchError("invalid frame length: 148 (max 40)"))
				})

			})
		})

		Context("with an empty buffer", func() {
			It("returns an error", func() {
				byteBuf := bytes.NewBuffer(nil)
				d := xivnet.NewDecoder(byteBuf, 32768)
				_, err := d.NextFrame()
				Expect(err).To(MatchError("peeking header failed reading 28 bytes from buffer: EOF"))
			})
		})

		Context("with an invalid header at the head of the buffer", func() {
			It("returns an error", func() {
				byteBuf := bytes.NewBuffer(invalidHeaderPacket)
				d := xivnet.NewDecoder(byteBuf, 32768)
				_, err := d.NextFrame()
				Expect(err).To(MatchError("invalid header: 52520000ff5d46e27f2a644d7b99c475e6f693da590100008a000000"))
			})
		})

		Context("with corrupt data on the buffer", func() {
			It("reads the first non-corrupted frame on the buffer", func() {
				byteBuf := bytes.NewBuffer(nil)
				_, err := byteBuf.Write(zlibPacket[:40])
				Expect(err).ToNot(HaveOccurred())
				_, err = byteBuf.Write(zlibPacket)
				Expect(err).ToNot(HaveOccurred())

				d := xivnet.NewDecoder(byteBuf, 32768)
				_, err = d.NextFrame()
				Expect(err.Error()).To(ContainSubstring("error decompressing data: zlib: invalid header"))

				f, err := d.NextFrame()
				Expect(err).ToNot(HaveOccurred())

				Expect(f).To(matchExpectedFrame(expectedZlibFrame))
			})
		})

		Context("with a block that specifies an invalid length", func() {
			It("returns an error", func() {
				byteBuf := bytes.NewBuffer(invalidBlockPacket)
				d := xivnet.NewDecoder(byteBuf, 32768)
				_, err := d.NextFrame()
				Expect(err).To(MatchError(
					"error decoding frame: error decoding blocks: not enough data: expected 32 bytes, got 24\n" +
						"Data: 000000000000000000000000000000000000000000000000400000000000010000000000000000002000000000000000000000000800000015cd5b0742e08958",
				))
			})
		})
	})

	Describe("CheckHeader", func() {
		Context("with an empty buffer", func() {
			It("returns an error", func() {
				byteBuf := bytes.NewBuffer(nil)
				d := xivnet.NewDecoder(byteBuf, 32768)
				_, err := d.CheckHeader()
				Expect(err).To(MatchError("peeking header failed reading 28 bytes from buffer: EOF"))
			})
		})

		Context("with an invalid header at the head of the buffer", func() {
			It("returns an error", func() {
				byteBuf := bytes.NewBuffer(invalidHeaderPacket)
				d := xivnet.NewDecoder(byteBuf, 32768)
				_, err := d.CheckHeader()
				Expect(err).To(MatchError("invalid header: 52520000ff5d46e27f2a644d7b99c475e6f693da590100008a000000"))
			})
		})
	})

	Describe("DiscardInvalidData", func() {
		Context("with an empty buffer", func() {
			It("does nothing to the buffer", func() {
				byteBuf := bytes.NewBuffer(nil)
				d := xivnet.NewDecoder(byteBuf, 32768)
				d.DiscardDataUntilValid()
				Expect(byteBuf.Len()).To(Equal(0))
			})
		})

		Context("with an invalid header at the head of the buffer", func() {
			It("discards the invalid data and allows the next decode operation to succeed with valid data", func() {
				byteBuf := bytes.NewBuffer(append(invalidHeaderPacket, zeroBlockPacket...))
				d := xivnet.NewDecoder(byteBuf, 32768)
				d.DiscardDataUntilValid()
				f, err := d.NextFrame()
				Expect(err).ToNot(HaveOccurred())
				Expect(f.Length).To(Equal(uint32(48)))
			})
		})

		It("does nothing when the header is already valid", func() {
			byteBuf := bytes.NewBuffer(zeroBlockPacket)
			d := xivnet.NewDecoder(byteBuf, 32768)
			d.DiscardDataUntilValid()
			f, err := d.NextFrame()
			Expect(err).ToNot(HaveOccurred())
			Expect(f.Length).To(Equal(uint32(48)))
		})
	})
})
