package sbvj

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/hollowness-inside/SBVJ01-Reader/pkg/types"
)

func ReadBytes(buf []byte) (*SBVJ, error) {
	buffer := bytes.NewBuffer(buf)
	return Read(buffer)
}

func ReadFile(path string) (*SBVJ, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return Read(file)
}

func Read(r io.Reader) (*SBVJ, error) {
	reader := bufio.NewReader(r)

	magic := make([]byte, 6)
	if _, err := io.ReadFull(reader, magic); err != nil {
		return nil, err
	}

	if !bytes.Equal(magic, []byte("SBVJ01")) {
		return nil, fmt.Errorf("wrong magic - expected SBVJ01, received %s", string(magic))
	}

	sbvj := SBVJ{}

	name, err := readString(reader)
	if err != nil {
		return nil, err
	}
	sbvj.Options.Name = name

	versioned, err := readBoolean(reader)
	if err != nil {
		return nil, err
	}
	sbvj.Options.Versioned = versioned

	if versioned {
		if err := binary.Read(reader, binary.BigEndian, &sbvj.Options.Version); err != nil {
			return nil, err
		}
	}

	object, err := readObject(reader)
	if err != nil {
		return nil, err
	}
	sbvj.Content = object

	return &sbvj, nil
}

func readByte(r *bufio.Reader) (byte, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, err
	}

	return b, nil
}

func readBytes(r *bufio.Reader) ([]byte, error) {
	size, err := readVarint(r)
	if err != nil {
		return nil, err
	}

	bytes := make([]byte, size)

	if _, err = io.ReadFull(r, bytes); err != nil {
		return nil, err
	}

	return bytes, nil
}

func readString(r *bufio.Reader) (string, error) {
	bytes, err := readBytes(r)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func readVarint(r *bufio.Reader) (int64, error) {
	var value int64
	for {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}

		if b&0b10000000 == 0 {
			value = value<<7 | int64(b)
			return value, nil
		}
		value = value<<7 | (int64(b) & 0b01111111)
	}
}

func readUVarint(r *bufio.Reader) (int64, error) {
	v, err := readVarint(r)
	if err != nil {
		return 0, err
	}

	if v&1 != 0 {
		return -(v >> 1) - 1, nil
	}

	return v >> 1, nil
}

func readDouble(r *bufio.Reader) (float64, error) {
	var val float64
	if err := binary.Read(r, binary.LittleEndian, &val); err != nil {
		return 0, nil
	}

	return val, nil
}

func readBoolean(r *bufio.Reader) (bool, error) {
	b, err := readByte(r)
	if err != nil {
		return false, nil
	}

	if b == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func readObject(r *bufio.Reader) (types.SBVJObject, error) {
	object := types.SBVJObject{}

	typ, err := readByte(r)
	if err != nil {
		return object, err
	}
	object.Type = types.SBVJType(typ)

	var value any
	switch object.Type {
	case types.NIL:
		object.Value = nil
	case types.DOUBLE:
		value, err = readDouble(r)
	case types.BOOLEAN:
		value, err = readBoolean(r)
	case types.VARINT:
		value, err = readUVarint(r)
	case types.STRING:
		value, err = readString(r)
	case types.LIST:
		value, err = readList(r)
	case types.MAP:
		value, err = readMap(r)
	default:
		return object, fmt.Errorf("unknown type: %d", object.Type)
	}

	if err != nil {
		return object, err
	}
	object.Value = value

	return object, nil
}

func readList(r *bufio.Reader) (types.SBVJList, error) {
	size, err := readVarint(r)
	if err != nil {
		return nil, err
	}

	sbvjList := make(types.SBVJList, size)

	var i int64
	for i = 0; i < size; i++ {
		token, err := readObject(r)
		if err != nil {
			return nil, err
		}

		sbvjList[i] = token
	}

	return sbvjList, nil
}

func readMap(r *bufio.Reader) (types.SBVJMap, error) {
	size, err := readVarint(r)
	if err != nil {
		return nil, err
	}

	sbvjmap := make(types.SBVJMap, size)

	var i int64
	for i = 0; i < size; i++ {
		key, err := readString(r)
		if err != nil {
			return nil, err
		}

		value, err := readObject(r)
		if err != nil {
			return nil, err
		}

		sbvjmap[key] = value
	}

	return sbvjmap, nil
}
