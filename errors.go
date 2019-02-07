package xivnet

import (
	"fmt"
)

// EOFError is an error that is returned when there is not enough
// data to process the packet
type EOFError struct {
	operation       string
	attemptedLength int
	wrapped         error
}

func (e EOFError) Error() string {
	return fmt.Sprintf(
		"%s failed reading %d bytes from buffer: %s",
		e.operation,
		e.attemptedLength,
		e.wrapped,
	)
}

// InvalidHeaderError is an error that is returned when the frame header is not
// something that is recognized by the decoder.
// It uses a string as a field to guarantee that the header is copied, so that
// changes to the original buffer don't affect the error.
type InvalidHeaderError struct {
	header string
}

func (e InvalidHeaderError) Error() string {
	return fmt.Sprintf("invalid header: %s", e.header)
}

// InvalidFrameLength is an error that is returned when the frame specifies
// a length that is too large. This error guards against the case when garbage
// or malicious data is read as part of decoding.
type InvalidFrameLengthError struct {
	length    uint32
	maxLength int
}

func (e InvalidFrameLengthError) Error() string {
	return fmt.Sprintf("invalid frame length: %d (max %d)", e.length, e.maxLength)
}

// MismatchedReadError is returned when somehow the amount of bytes read
// doesn't match the requested length. We don't expect this error to happen
// because we have previously already peeked the required amount of data.
type MismatchedReadLengthsError struct {
	readLength     int
	expectedLength int
}

func (e MismatchedReadLengthsError) Error() string {
	return fmt.Sprintf(
		"mismatched read lengths: %d != %d", e.readLength, e.expectedLength,
	)
}

// DecodingError is returned whenenever some error occurs while decoding the
// frame or some block within the frame.
type DecodingError struct {
	wrapped   error
	debugData string
}

func (e DecodingError) Error() string {
	return fmt.Sprintf(
		"error decoding frame: %s\nData: %s", e.wrapped, e.debugData,
	)
}
