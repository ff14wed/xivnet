package xivnet_test

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"

	"github.com/ff14wed/xivnet"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Decoder", func() {
	var logger *log.Logger
	BeforeEach(func() {
		logger = log.New(ioutil.Discard, "", log.LstdFlags)
	})
	Describe("Decode", func() {
		Context("with zlib compressed blocks", func() {
			It("properly decodes a packet into the correct structures", func() {
				buf := bufio.NewReader(bytes.NewReader(zlibPacket))
				d := xivnet.NewDecoder(32768, logger)
				f, err := d.Decode(buf)
				Expect(err).ToNot(HaveOccurred())
				Expect(f.Header).To(Equal(expectedZlibFrame.Header))
				Expect(f.Time).To(Equal(expectedZlibFrame.Time))
				Expect(f.Length).To(Equal(expectedZlibFrame.Length))
				Expect(f.NumBlocks).To(Equal(expectedZlibFrame.NumBlocks))
				Expect(f.Compression).To(Equal(expectedZlibFrame.Compression))
				Expect(f.Blocks).To(Equal(expectedZlibFrame.Blocks))

				Expect(f.Reserved1).To(Equal(uint16(0)))
				Expect(f.Reserved2).To(Equal(uint32(0)))
				Expect(f.Reserved3).To(Equal(uint16(0)))
			})
		})

		Context("with a zero-block zlib compressed frame", func() {
			It("correct returns a frame with no blocks", func() {
				byteBuf := bytes.NewBuffer(zeroBlockPacket)
				buf := bufio.NewReader(byteBuf)
				d := xivnet.NewDecoder(32768, logger)
				f, err := d.Decode(buf)
				Expect(err).ToNot(HaveOccurred())
				Expect(f.NumBlocks).To(BeZero())
				Expect(f.Blocks).To(BeEmpty())
			})
		})
		Context("with multiple frames on the buffer", func() {
			It("properly decodes the packets into the correct structures", func() {
				buf := bufio.NewReader(bytes.NewBuffer(append(zlibPacket, zlibPacket...)))

				d := xivnet.NewDecoder(32768, logger)
				for i := 0; i < 2; i++ {
					f, err := d.Decode(buf)
					Expect(err).ToNot(HaveOccurred())
					Expect(f.Header).To(Equal(expectedZlibFrame.Header))
					Expect(f.Time).To(Equal(expectedZlibFrame.Time))
					Expect(f.Length).To(Equal(expectedZlibFrame.Length))
					Expect(f.NumBlocks).To(Equal(expectedZlibFrame.NumBlocks))
					Expect(f.Compression).To(Equal(expectedZlibFrame.Compression))
					Expect(f.Blocks).To(Equal(expectedZlibFrame.Blocks))

					Expect(f.Reserved1).To(Equal(uint16(0)))
					Expect(f.Reserved2).To(Equal(uint32(0)))
					Expect(f.Reserved3).To(Equal(uint16(0)))
				}
			})
		})

		Context("with a non-zlib compressed packet and short block data", func() {
			It("properly decodes a packet into the correct structures", func() {
				buf := bufio.NewReader(bytes.NewBuffer(nonZlibPacket))
				d := xivnet.NewDecoder(32768, logger)
				frame, err := d.Decode(buf)
				Expect(err).ToNot(HaveOccurred())
				Expect(frame.Header).To(Equal(expectedNonZlibFrame.Header))
				Expect(frame.Time).To(Equal(expectedNonZlibFrame.Time))
				Expect(frame.Length).To(Equal(expectedNonZlibFrame.Length))
				Expect(frame.NumBlocks).To(Equal(expectedNonZlibFrame.NumBlocks))
				Expect(frame.Compression).To(Equal(expectedNonZlibFrame.Compression))
				Expect(frame.Blocks).To(Equal(expectedNonZlibFrame.Blocks))

				Expect(frame.Reserved1).To(Equal(uint16(0)))
				Expect(frame.Reserved2).To(Equal(uint32(0)))
				Expect(frame.Reserved3).To(Equal(uint16(0)))
			})
		})

		Context("with incomplete data on the buffer", func() {
			It("eventually is able to read the frame", func() {
				byteBuf := bytes.NewBuffer(nil)
				_, err := byteBuf.Write(zlibPacket[:69])
				Expect(err).ToNot(HaveOccurred())
				buf := bufio.NewReader(byteBuf)

				d := xivnet.NewDecoder(32768, logger)
				_, err = d.Decode(buf)
				Expect(err).To(MatchError(xivnet.ErrNotEnoughData))

				_, err = byteBuf.Write(zlibPacket[69:])
				Expect(err).ToNot(HaveOccurred())

				f, err := d.Decode(buf)
				Expect(err).ToNot(HaveOccurred())
				Expect(f.Header).To(Equal(expectedZlibFrame.Header))
				Expect(f.Time).To(Equal(expectedZlibFrame.Time))
				Expect(f.Length).To(Equal(expectedZlibFrame.Length))
				Expect(f.NumBlocks).To(Equal(expectedZlibFrame.NumBlocks))
				Expect(f.Compression).To(Equal(expectedZlibFrame.Compression))
				Expect(f.Blocks).To(Equal(expectedZlibFrame.Blocks))

				Expect(f.Reserved1).To(Equal(uint16(0)))
				Expect(f.Reserved2).To(Equal(uint32(0)))
				Expect(f.Reserved3).To(Equal(uint16(0)))
			})
		})

		Context("with a decoder that has too small a buffer", func() {
			It("returns an error", func() {
				buf := bufio.NewReader(bytes.NewReader(zlibPacket))
				d := xivnet.NewDecoder(8, logger)
				_, err := d.Decode(buf)
				Expect(err).To(MatchError("frame is too large: 148 > 8"))
			})
		})

		Context("with an empty buffer", func() {
			It("returns an error", func() {
				byteBuf := bytes.NewBuffer(nil)
				buf := bufio.NewReader(byteBuf)
				d := xivnet.NewDecoder(32768, logger)
				_, err := d.Decode(buf)
				Expect(err).To(MatchError(xivnet.ErrNotEnoughData))
			})
		})

		Context("with an invalid header at the head of the buffer", func() {
			It("returns an error", func() {
				byteBuf := bytes.NewBuffer(invalidHeaderPacket)
				buf := bufio.NewReader(byteBuf)
				d := xivnet.NewDecoder(32768, logger)
				_, err := d.Decode(buf)
				Expect(err).To(MatchError(xivnet.ErrInvalidHeader))
			})
		})
	})
	Describe("CheckHeader", func() {
		Context("with an empty buffer", func() {
			It("returns an error", func() {
				byteBuf := bytes.NewBuffer(nil)
				buf := bufio.NewReader(byteBuf)
				d := xivnet.NewDecoder(32768, logger)
				_, err := d.CheckHeader(buf)
				Expect(err).To(MatchError(xivnet.ErrNotEnoughData))
			})
		})

		Context("with an invalid header at the head of the buffer", func() {
			It("returns an error", func() {
				byteBuf := bytes.NewBuffer(invalidHeaderPacket)
				buf := bufio.NewReader(byteBuf)
				d := xivnet.NewDecoder(32768, logger)
				_, err := d.CheckHeader(buf)
				Expect(err).To(MatchError(xivnet.ErrInvalidHeader))
			})
		})
	})
	Describe("DiscardInvalidData", func() {
		Context("with an empty buffer", func() {
			It("does nothing to the buffer", func() {
				byteBuf := bytes.NewBuffer(nil)
				buf := bufio.NewReader(byteBuf)
				d := xivnet.NewDecoder(32768, logger)
				d.DiscardDataUntilValid(buf)
				Expect(byteBuf.Len()).To(Equal(0))
			})
		})

		Context("with an invalid header at the head of the buffer", func() {
			It("discards the invalid data and allows the next decode operation to succeed with valid data", func() {
				byteBuf := bytes.NewBuffer(append(invalidHeaderPacket, zeroBlockPacket...))
				buf := bufio.NewReader(byteBuf)
				d := xivnet.NewDecoder(32768, logger)
				d.DiscardDataUntilValid(buf)
				f, err := d.Decode(buf)
				Expect(err).ToNot(HaveOccurred())
				Expect(f.Length).To(Equal(uint32(48)))
			})
		})
	})
})
