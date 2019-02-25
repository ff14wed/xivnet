package bytestring

import (
	"bytes"
	"encoding/hex"
	"strconv"
	"strings"
	"sync"
)

var mbufPool *marshalBufferPool

type marshalBufferPool struct {
	pool *sync.Pool
}

func newPool() *marshalBufferPool {
	return &marshalBufferPool{
		pool: new(sync.Pool),
	}
}

func (m *marshalBufferPool) Get() *bytes.Buffer {
	b := m.pool.Get()
	if b == nil {
		return new(bytes.Buffer)
	}
	return b.(*bytes.Buffer)
}

func (m *marshalBufferPool) Put(b *bytes.Buffer) {
	b.Reset()
	m.pool.Put(b)
}

func BytesToByteString(b []byte) ([]byte, error) {
	n := len(b)
	if n == 0 {
		return []byte(`""`), nil
	}
	if mbufPool == nil {
		mbufPool = newPool()
	}
	marshalBuffer := mbufPool.Get()
	defer mbufPool.Put(marshalBuffer)

	marshalBuffer.WriteRune('"')
	if b[0] < 0x10 {
		marshalBuffer.WriteRune('0')
	}
	marshalBuffer.WriteString(strconv.FormatUint(uint64(b[0]), 16))
	for i := 1; i < n; i++ {
		marshalBuffer.WriteRune(' ')
		if b[i] < 0x10 {
			marshalBuffer.WriteRune('0')
		}
		marshalBuffer.WriteString(strconv.FormatUint(uint64(b[i]), 16))
	}
	marshalBuffer.WriteRune('"')
	return marshalBuffer.Bytes(), nil
}

func ByteStringToBytes(b []byte) ([]byte, error) {
	n := len(b)
	encodedHex := strings.Replace(string(b[1:n-1]), " ", "", -1)
	return hex.DecodeString(encodedHex)
}
