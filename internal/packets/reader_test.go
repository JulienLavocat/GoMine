package packets_test

import (
	"testing"

	"github.com/JulienLavocat/gomine/internal/packets"
)

func TestReader_Bool(t *testing.T) {
	type test struct {
		input  []byte
		expect []bool
	}

	tests := map[string]test{
		"read true":                 {input: []byte{0x01}, expect: []bool{true}},
		"read false":                {input: []byte{0x00}, expect: []bool{false}},
		"read false multiple bytes": {input: []byte{0x00, 0x01}, expect: []bool{false}},
		"read true multiple bytes":  {input: []byte{0x01, 0x00}, expect: []bool{true}},
		"read increment position":   {input: []byte{0x01, 0x00, 0x01, 0x01, 0x00, 0x00}, expect: []bool{true, false, true, true, false, false}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			reader := packets.NewReader(tc.input)

			for index, expect := range tc.expect {
				var result bool
				err := reader.Read(&result)
				if err != nil {
					t.Error(err)
					continue
				}
				if result != expect {
					t.Errorf("expected %v but got %v at expect %v", tc.expect, result, index)
				}
			}
		})
	}
}

func TestReader_UInt8(t *testing.T) {
	type test struct {
		value []byte
	}

	tests := map[string]test{
		"read single uint8":   {value: []byte{123}},
		"read multiple uint8": {value: []byte{0, 128, 129, 255}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			reader := packets.NewReader(tc.value)

			for _, value := range tc.value {
				var result uint8
				err := reader.Read(&result)
				if err != nil {
					t.Error(err)
					continue
				}

				if result != value {
					t.Errorf("expected: %v got: %v", value, result)
				}
			}
		})
	}
}

func TestReader_Int8(t *testing.T) {
	type test struct {
		value  []byte
		expect []int8
	}

	tests := map[string]test{
		"read single positive int8": {value: []byte{123}, expect: []int8{123}},
		"read single negative int8": {value: []byte{133}, expect: []int8{-123}},
		"read multiple int8":        {value: []byte{128, 255, 0, 1, 127}, expect: []int8{-128, -1, 0, 1, 127}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			reader := packets.NewReader(tc.value)

			for _, value := range tc.expect {
				var result int8
				err := reader.Read(&result)
				if err != nil {
					t.Error(err)
					continue
				}

				if result != value {
					t.Errorf("expected: %v got: %v", value, result)
				}
			}
		})
	}
}

func TestReader_Int16(t *testing.T) {
	type test struct {
		value  []byte
		expect []int16
	}

	tests := map[string]test{
		"read single positive int16": {value: []byte{127, 255}, expect: []int16{32767}},
		"read single negative int16": {value: []byte{128, 0}, expect: []int16{-32768}},
		"read multiple int16":        {value: []byte{128, 0, 255, 255, 0, 0, 0, 1, 127, 255}, expect: []int16{-32768, -1, 0, 1, 32767}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			reader := packets.NewReader(tc.value)

			for _, value := range tc.expect {
				var result int16
				err := reader.Read(&result)
				if err != nil {
					t.Error(err)
					continue
				}

				if result != value {
					t.Errorf("expected: %v got: %v", value, result)
				}
			}
		})
	}
}
