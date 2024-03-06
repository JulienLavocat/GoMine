package packets

import (
	"bytes"
	"encoding/binary"
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
	return binary.Read(r.data, binary.BigEndian, data)
}

func (r *Reader) Bool() (bool, error) {
	var value bool
	err := binary.Read(r.data, binary.BigEndian, &value)
	if err != nil {
		return false, err
	}
	return value, nil
}

func (r *Reader) Int8() (int8, error) {
	var value int8
	err := binary.Read(r.data, binary.BigEndian, &value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (r *Reader) UInt8() (uint8, error) {
	var value uint8
	err := binary.Read(r.data, binary.BigEndian, &value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (r *Reader) Int16() (int16, error) {
	var value int16
	err := binary.Read(r.data, binary.BigEndian, &value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (r *Reader) UInt16() (uint16, error) {
	var value uint16
	err := binary.Read(r.data, binary.BigEndian, &value)
	if err != nil {
		return 0, err
	}
	return value, nil
}
