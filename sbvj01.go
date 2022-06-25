package sbvj01

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

type SBVJ01Type byte

const (
	NIL     SBVJ01Type = 0x01
	DOUBLE  SBVJ01Type = 0x02
	BOOLEAN SBVJ01Type = 0x03
	VARINT  SBVJ01Type = 0x04
	STRING  SBVJ01Type = 0x05
	LIST    SBVJ01Type = 0x06
	MAP     SBVJ01Type = 0x07
)

type SBVJ01 struct {
	Name      string
	Versioned bool
	Version   int32
	Value     SBVJ01Token
}

func ReadSBVJ01File(path string) SBVJ01 {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	magic := make([]byte, 6)
	n, err := reader.Read(magic)
	if err != nil {
		panic(err)
	}
	if n != 6 {
		panic("Cannot read magic")
	}
	if string(magic) != "SBVJ01" {
		panic("Wrong magic")
	}

	sbvj := SBVJ01{}
	size := readByte(reader)
	buffer := make([]byte, size)
	n, err = reader.Read(buffer)
	if n != int(size) || err != nil {
		panic(err)
	}
	sbvj.Name = string(buffer)

	if readByte(reader) == 0 {
		sbvj.Versioned = false
	} else {
		sbvj.Versioned = true
		err := binary.Read(reader, binary.BigEndian, &sbvj.Version)
		if err != nil {
			panic(err)
		}
	}

	sbvj.Value = readToken(reader)
	return sbvj
}

func readByte(r *bufio.Reader) byte {
	b, err := r.ReadByte()
	if err != nil {
		panic(err)
	}

	return b
}

func readString(r *bufio.Reader) string {
	bytes := readBytes(r)
	return string(bytes)
}

func readBytes(r *bufio.Reader) []byte {
	size := readVarint(r)

	bytes := make([]byte, size)
	var i int64 = 0
	for i < size {
		b, err := r.ReadByte()
		if err != nil {
			panic(err)
		}

		bytes[i] = b
		i += 1
	}

	return bytes
}

func readVarint(r *bufio.Reader) int64 {
	var value int64
	for {
		b, err := r.ReadByte()
		if err != nil {
			panic(err)
		}

		if b&0b10000000 == 0 {
			value = value<<7 | int64(b)
			return value
		}
		value = value<<7 | (int64(b) & 0b01111111)
	}
}

func readSignedVarint(r *bufio.Reader) int64 {
	v := readVarint(r)

	if v&1 != 0 {
		return -(v >> 1) - 1
	}

	return v >> 1
}

func readToken(r *bufio.Reader) SBVJ01Token {
	token := SBVJ01Token{}

	tp := readByte(r)
	token.Type = SBVJ01Type(tp)

	switch token.Type {
	case NIL:
		token.Value = nil
	case DOUBLE:
		token.Value = readDouble(r)
	case BOOLEAN:
		token.Value = readBoolean(r)
	case VARINT:
		token.Value = readSignedVarint(r)
	case STRING:
		token.Value = readString(r)
	case LIST:
		token.Value = readList(r)
	case MAP:
		token.Value = readMap(r)
	default:
		panic(fmt.Sprintf("Unknown token type %d", token.Type))
	}

	return token
}

func readDouble(r *bufio.Reader) float64 {
	var val float64
	err := binary.Read(r, binary.LittleEndian, &val)
	if err != nil {
		panic(err)
	}

	return val
}

func readBoolean(r *bufio.Reader) bool {
	if readByte(r) == 0 {
		return false
	} else {
		return true
	}
}

func readList(r *bufio.Reader) SBVJ01List {
	sbvjList := SBVJ01List{}

	size := readVarint(r)
	sbvjList.Items = make([]SBVJ01Token, size)

	var i int64
	for i = 0; i < size; i++ {
		token := readToken(r)
		sbvjList.Items[i] = token
	}

	return sbvjList
}

func readMap(r *bufio.Reader) SBVJ01Map {
	sbvjmap := SBVJ01Map{}

	size := readVarint(r)
	sbvjmap.Items = make([]SBVJ01Pair, size)

	var i int64
	for i = 0; i < size; i++ {
		key := readString(r)
		value := readToken(r)
		sbvjmap.Items[i] = SBVJ01Pair{key, value}
	}

	return sbvjmap
}
