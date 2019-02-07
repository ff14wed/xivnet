package xivnet_test

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"time"

	"github.com/ff14wed/xivnet/v2"
	"github.com/ff14wed/xivnet/v2/datatypes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Encoder", func() {
	Describe("Frame", func() {
		var (
			incorrectFrame   xivnet.Frame
			realBlockLengths []uint32
		)

		BeforeEach(func() {
			incorrectFrame = expectedZlibFrame
			incorrectFrame.Blocks = nil
			incorrectFrame.Length = 0

			for _, b := range expectedZlibFrame.Blocks {
				var newB xivnet.Block = *b
				realBlockLengths = append(realBlockLengths, newB.Length)
				newB.Length = 0
				incorrectFrame.Blocks = append(incorrectFrame.Blocks, &newB)
			}
		})

		Describe("CorrectLength", func() {
			Context("with compression", func() {
				It("sets the correct length of the frame and its blocks", func() {
					incorrectFrame.CorrectLength()
					err := incorrectFrame.CompressBlocks()
					Expect(err).ToNot(HaveOccurred())

					var totalBlockLength uint32
					for i, b := range incorrectFrame.Blocks {
						Expect(b.Length).To(Equal(realBlockLengths[i]))
						totalBlockLength += b.Length
					}

					Expect(incorrectFrame.Length).To(BeNumerically("<", 40+totalBlockLength))
				})
			})
			Context("without compression", func() {
				It("sets the correct length of the frame and its blocks", func() {
					incorrectFrame.CorrectLength()
					var totalBlockLength uint32
					for i, b := range incorrectFrame.Blocks {
						Expect(b.Length).To(Equal(realBlockLengths[i]))
						totalBlockLength += b.Length
					}
					Expect(incorrectFrame.Length).To(Equal(40 + totalBlockLength))
				})
			})

		})
		Describe("CorrectTimestamps", func() {
			It("sets the time to the passed in argument on the frame", func() {
				target := time.Now().Add(-3 * time.Second)
				incorrectFrame.CorrectTimestamps(target)
				Expect(incorrectFrame.Time).To(BeTemporally("==", target))
				for _, b := range incorrectFrame.Blocks {
					Expect(b.Header.Time).To(BeTemporally("==", target))
				}
			})
		})

		Describe("Encode", func() {
			var (
				decoder *xivnet.Decoder
			)
			BeforeEach(func() {
				decoder = xivnet.NewDecoder(32768)
			})
			Context("with zlib compression", func() {
				It("writes the correct encoding of the frame to the writer", func() {
					buf := new(bytes.Buffer)
					err := incorrectFrame.Encode(buf, expectedZlibFrame.Time, true)
					Expect(err).ToNot(HaveOccurred())

					Expect(buf.Bytes()).To(Equal(zlibPacket))
				})
			})
			Context("without zlib compression", func() {
				It("writes the correct encoding of the frame to the writer", func() {
					buf := new(bytes.Buffer)
					err := incorrectFrame.Encode(buf, expectedZlibFrame.Time, false)
					Expect(err).ToNot(HaveOccurred())

					// Test that the raw SubjectID is there
					Expect(buf.Bytes()[44:48]).To(Equal([]byte{0x15, 0xCD, 0x5B, 0x07}))

					decoderBuf := bufio.NewReader(buf)
					decodedFrame, err := decoder.Decode(decoderBuf)
					Expect(err).ToNot(HaveOccurred())

					Expect(decodedFrame.Header).To(Equal(expectedZlibFrame.Header))
					Expect(decodedFrame.Time).To(Equal(expectedZlibFrame.Time))
					Expect(decodedFrame.NumBlocks).To(Equal(expectedZlibFrame.NumBlocks))
					Expect(decodedFrame.Compression).To(Equal(uint16(1)))
					Expect(decodedFrame.Blocks).To(Equal(expectedZlibFrame.Blocks))
					Expect(decodedFrame.Reserved1).To(Equal(expectedZlibFrame.Reserved1))
					Expect(decodedFrame.Reserved2).To(Equal(expectedZlibFrame.Reserved2))
					Expect(decodedFrame.Reserved3).To(Equal(expectedZlibFrame.Reserved3))
				})
			})
		})
	})

	Describe("Block", func() {
		var (
			incorrectBlock  xivnet.Block
			realBlockLength uint32
		)
		BeforeEach(func() {
			incorrectBlock = *expectedZlibFrame.Blocks[0]
			realBlockLength = incorrectBlock.Length
			incorrectBlock.Length = 0
		})
		Describe("CorrectLength", func() {
			It("sets the correct length of the block", func() {
				incorrectBlock.CorrectLength()
				Expect(incorrectBlock.Length).To(Equal(realBlockLength))
			})
		})

		Describe("Encode", func() {
			It("writes the correct encoding of the block to the writer", func() {
				buf := new(bytes.Buffer)
				err := expectedZlibFrame.Blocks[0].Encode(buf)
				Expect(err).ToNot(HaveOccurred())
				Expect(buf.Bytes()).To(Equal(bytesZlibBlock0))
			})
			Context("with typed block data", func() {
				var (
					movement           *datatypes.Movement
					expectedBlockBytes []byte
				)
				BeforeEach(func() {
					movement = &datatypes.Movement{
						Direction: 0x12,
						U1:        0x12,
						U2:        0x4567,
						Position:  datatypes.PackedPosition{X: 0x89AB, Z: 0x89AB, Y: 0x89AB},
						U3:        0x4567,
					}
					expectedBlockBytes = []byte{
						0x30, 0x00, 0x00, 0x00, // Length
						0x15, 0xCD, 0x5B, 0x07, // SubjectID
						0x15, 0xCD, 0x5B, 0x07, // CurrentID
						0x0C, 0x00, 0x00, 0x00, // U1
						0x14, 0x00, 0x42, 0x01, // U2 and Opcode
						0x22, 0x00, 0x00, 0x00, // U3
						0x3f, 0xe0, 0x89, 0x58, // Time
						0x00, 0x00, 0x00, 0x00, // U4
						// Block Data begins here
						0x12, 0x12, 0x67, 0x45, 0x00, 0x00, // Direction, U1, U2
						0xAB, 0x89, 0xAB, 0x89, 0xAB, 0x89, // PackedPosition
						0x67, 0x45, 0x00, 0x00, // U3
					}
				})
				It("marshals typed block data to the writer", func() {
					buf := new(bytes.Buffer)
					testBlock := *expectedZlibFrame.Blocks[0]
					testBlock.Data = movement
					testBlock.Length = uint32(32 + binary.Size(movement))

					err := testBlock.Encode(buf)
					Expect(err).ToNot(HaveOccurred())
					Expect(buf.Bytes()).To(Equal(expectedBlockBytes))
				})
			})
			It("errors for blocks that are too small", func() {
				buf := new(bytes.Buffer)
				err := incorrectBlock.Encode(buf)
				Expect(err).To(MatchError("Block length is too small"))
			})
		})
	})
})
