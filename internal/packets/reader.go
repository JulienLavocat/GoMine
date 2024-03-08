package packets

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math"
)

type Reader struct {
	data     *bytes.Reader
	position int
	size     int
}

func NewReader(data []byte) *Reader {
	return &Reader{data: bytes.NewReader(data), size: len(data), position: 0}
}

func (r *Reader) Read(data any) error {
	// Code from standard library (binary.go#233)
	n := intDataSize(data)
	if n == 0 {
		return errors.New("unsupported type")
	}

	bs := make([]byte, n)
	if _, err := io.ReadFull(r.data, bs); err != nil {
		return err
	}
	switch data := data.(type) {
	case *bool:
		*data = bs[0] != 0
	case *int8:
		*data = int8(bs[0])
	case *uint8:
		*data = bs[0]
	case *int16:
		*data = int16(binary.BigEndian.Uint16(bs))
	case *uint16:
		*data = binary.BigEndian.Uint16(bs)
	case *int32:
		*data = int32(binary.BigEndian.Uint32(bs))
	case *uint32:
		*data = binary.BigEndian.Uint32(bs)
	case *int64:
		*data = int64(binary.BigEndian.Uint64(bs))
	case *uint64:
		*data = binary.BigEndian.Uint64(bs)
	case *float32:
		*data = math.Float32frombits(binary.BigEndian.Uint32(bs))
	case *float64:
		*data = math.Float64frombits(binary.BigEndian.Uint64(bs))
	case []bool:
		for i, x := range bs { // Easier to loop over the input for 8-bit values.
			data[i] = x != 0
		}
	case []int8:
		for i, x := range bs {
			data[i] = int8(x)
		}
	case []uint8:
		copy(data, bs)
	case []int16:
		for i := range data {
			data[i] = int16(binary.BigEndian.Uint16(bs[2*i:]))
		}
	case []uint16:
		for i := range data {
			data[i] = binary.BigEndian.Uint16(bs[2*i:])
		}
	case []int32:
		for i := range data {
			data[i] = int32(binary.BigEndian.Uint32(bs[4*i:]))
		}
	case []uint32:
		for i := range data {
			data[i] = binary.BigEndian.Uint32(bs[4*i:])
		}
	case []int64:
		for i := range data {
			data[i] = int64(binary.BigEndian.Uint64(bs[8*i:]))
		}
	case []uint64:
		for i := range data {
			data[i] = binary.BigEndian.Uint64(bs[8*i:])
		}
	case []float32:
		for i := range data {
			data[i] = math.Float32frombits(binary.BigEndian.Uint32(bs[4*i:]))
		}
	case []float64:
		for i := range data {
			data[i] = math.Float64frombits(binary.BigEndian.Uint64(bs[8*i:]))
		}
	default:
		n = 0 // fast path doesn't apply
	}

	return nil
}

// intDataSize returns the size of the data required to represent the data when encoded.
// It returns zero if the type cannot be implemented by the fast path in Read or Write.
func intDataSize(data any) int {
	switch data := data.(type) {
	case bool, int8, uint8, *bool, *int8, *uint8:
		return 1
	case []bool:
		return len(data)
	case []int8:
		return len(data)
	case []uint8:
		return len(data)
	case int16, uint16, *int16, *uint16:
		return 2
	case []int16:
		return 2 * len(data)
	case []uint16:
		return 2 * len(data)
	case int32, uint32, *int32, *uint32:
		return 4
	case []int32:
		return 4 * len(data)
	case []uint32:
		return 4 * len(data)
	case int64, uint64, *int64, *uint64:
		return 8
	case []int64:
		return 8 * len(data)
	case []uint64:
		return 8 * len(data)
	case float32, *float32:
		return 4
	case float64, *float64:
		return 8
	case []float32:
		return 4 * len(data)
	case []float64:
		return 8 * len(data)
	}
	return 0
}
