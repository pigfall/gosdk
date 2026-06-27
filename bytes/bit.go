package bytes

import (
	"encoding/binary"
	"math"
)

func HowManyBitIsOneInByte(bte byte) int {
	var count int
	for i := 7; i >= 0; i-- {
		if ((bte >> i) & (1)) > 0 {
			count++
		}
	}
	return count
}

func Float32SliceToBytes(values []float32, order binary.ByteOrder) []byte {
	bytes := make([]byte, len(values)*4)
	for i, v := range values {
		order.PutUint32(bytes[i*4:], math.Float32bits(v))
	}
	return bytes
}

func Uint32SliceToBytes(values []uint32, order binary.ByteOrder) []byte {
	b := make([]byte, len(values)*4)
	for i, v := range values {
		order.PutUint32(b[i*4:], v)
	}

	return b
}
