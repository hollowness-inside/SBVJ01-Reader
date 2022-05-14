package sbvj01

import (
	"bufio"
	"encoding/binary"
	"errors"
	"os"
)

const (
	NIL     = 0x01
	DOUBLE  = 0x02
	BOOLEAN = 0x03
	VARINT  = 0x04
	STRING  = 0x05
	LIST    = 0x06
	MAP     = 0x07
)

type SBVJ01 struct {
	Name      string
	Versioned bool
	Version   int32
	Value     SBVJ01Token
}

func ReadSBVJ01File(path string) (*SBVJ01, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	magic := make([]byte, 6)
	_, err = reader.Read(magic)
	if err != nil {
		return nil, err
	}

	if string(magic) != "SBVJ01" {
		return nil, errors.New("this is not a sbvj01 file")
	}

	sbvj := new(SBVJ01)
	sbvj.Name, err = readString(reader)
	if err != nil {
		return nil, err
	}

	versioned, _ := reader.ReadByte()
	if versioned == 0 {
		sbvj.Versioned = false
	} else {
		sbvj.Versioned = true
		err = binary.Read(reader, binary.LittleEndian, &sbvj.Version)
	}

	if err != nil {
		return nil, err
	}

	sbvj.Value, err = readToken(reader)
	if err != nil {
		return nil, err
	}

	return sbvj, nil
}

func readString(r *bufio.Reader) (string, error) {
	bytes, err := readBytes(r)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func readBytes(r *bufio.Reader) ([]byte, error) {
	size, err := readVarint(r)
	if err != nil {
		return nil, err
	}

	bytes := make([]byte, size)
	r.Read(bytes)

	return bytes, nil
}

func readVarint(r *bufio.Reader) (int, error) {
	v, err := binary.ReadUvarint(r)
	if err != nil {
		return 0, err
	}

	return int(v), nil
}

func readToken(r *bufio.Reader) (SBVJ01Token, error) {
	token := SBVJ01Token{}

	token.Type, _ = r.ReadByte()

	var value any
	var err error
	switch token.Type {
	case NIL:
		token.Value = nil
	case DOUBLE:
		value, err = readDouble(r)
	case BOOLEAN:
		value, err = readBoolean(r)
	case VARINT:
		value, err = readVarint(r)
	case STRING:
		value, err = readString(r)
	case LIST:
		value, err = readList(r)
	case MAP:
		value, err = readMap(r)
	}

	if err != nil {
		return token, err
	}

	token.Value = value

	return token, nil
}

func readDouble(r *bufio.Reader) (float64, error) {
	var val float64
	err := binary.Read(r, binary.LittleEndian, &val)
	if err != nil {
		return 0, err
	}

	return val, nil
}

func readBoolean(r *bufio.Reader) (bool, error) {
	b, err := r.ReadByte()
	if err != nil {
		return false, err
	}

	if b == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func readList(r *bufio.Reader) (SBVJ01List, error) {
	size, err := readVarint(r)
	if err != nil {
		return SBVJ01List{}, err
	}

	list := NewSBVJ01List(size)
	for i := 0; i < size; i++ {
		token, err := readToken(r)
		if err != nil {
			return list, err
		}

		list.Items[i] = token
	}

	return list, nil
}

func readMap(r *bufio.Reader) (SBVJ01Map, error) {
	size, err := readVarint(r)
	if err != nil {
		return SBVJ01Map{}, err
	}

	sbvjmap := NewSBVJ01Map(size)
	for i := 0; i < size; i++ {
		key, err := readString(r)
		if err != nil {
			return sbvjmap, err
		}

		value, err := readToken(r)
		if err != nil {
			return sbvjmap, err
		}

		sbvjmap.Items[i] = SBVJ01Pair{key, value}
	}

	return sbvjmap, nil
}
