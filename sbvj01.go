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
	Value     *SBVJ01Token
}

type SBVJ01Token struct {
	Type  byte
	Value any
}

type SBVJ01Pair struct {
	Key   string
	Value *SBVJ01Token
}

func ReadSBVJ01File(path string) (*SBVJ01, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	magic := make([]byte, 6)
	reader.Read(magic)

	if string(magic) != "SBVJ01" {
		return nil, errors.New("this is not a sbvj01 file")
	}

	sbvj := new(SBVJ01)

	sbvj.Name = readString(reader)

	versioned, _ := reader.ReadByte()
	if versioned == 0 {
		sbvj.Versioned = false
	} else {
		sbvj.Versioned = true
		binary.Read(reader, binary.LittleEndian, &sbvj.Version)
	}

	sbvj.Value = readToken(reader)

	return sbvj, nil
}

func readString(r *bufio.Reader) string {
	return string(readBytes(r))
}

func readBytes(r *bufio.Reader) []byte {
	size := readVarint(r)

	bytes := make([]byte, size)
	r.Read(bytes)

	return bytes
}

func readVarint(r *bufio.Reader) int {
	// TODO: Supposed to read VLQ, not a byte
	v, _ := r.ReadByte()
	return int(v)
}

func readToken(r *bufio.Reader) *SBVJ01Token {
	token := new(SBVJ01Token)

	token.Type, _ = r.ReadByte()

	switch token.Type {
	case NIL:
		token.Value = nil
	case DOUBLE:
		token.Value = readDouble(r)
	case BOOLEAN:
		token.Value = readBoolean(r)
	case VARINT:
		token.Value = readVarint(r)
	case STRING:
		token.Value = readString(r)
	case LIST:
		token.Value = readList(r)
	case MAP:
		token.Value = readMap(r)
	default:
		return nil
	}

	return token
}

func readDouble(r *bufio.Reader) float64 {
	var val float64
	binary.Read(r, binary.LittleEndian, &val)
	return val
}

func readBoolean(r *bufio.Reader) bool {
	b, _ := r.ReadByte()
	if b == 0 {
		return false
	} else {
		return true
	}
}

func readList(r *bufio.Reader) []*SBVJ01Token {
	size := readVarint(r)
	list := make([]*SBVJ01Token, size)
	for i := 0; i < size; i++ {
		list[i] = readToken(r)
	}

	return list
}

func readMap(r *bufio.Reader) []*SBVJ01Pair {
	size := readVarint(r)
	list := make([]*SBVJ01Pair, size)
	for i := 0; i < size; i++ {
		pair := new(SBVJ01Pair)
		pair.Key = readString(r)
		pair.Value = readToken(r)
		list[i] = pair
	}

	return list
}
