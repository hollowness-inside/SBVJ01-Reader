package sbvj01

import (
	"bufio"
	"encoding/binary"
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

func ReadSBVJ01File(path string) *SBVJ01 {
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
	sbvj.Name = readString(reader)

	versioned := readByte(reader)

	if versioned == 0 {
		sbvj.Versioned = false
	} else {
		sbvj.Versioned = true
		err = binary.Read(reader, binary.LittleEndian, &sbvj.Version)
	}

	sbvj.Value = readToken(reader)
	return &sbvj
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
	n, err := r.Read(bytes)
	if err != nil {
		panic(err)
	}

	if int64(n) != size {
		panic("Cannot read the needed amount of data")
	}

	return bytes
}

func readVarint(r *bufio.Reader) int64 {
	v, err := binary.ReadVarint(r)
	if err != nil {
		panic(err)
	}

	return v
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
		token.Value = readVarint(r)
	case STRING:
		token.Value = readString(r)
	case LIST:
		token.Value = readList(r)
	case MAP:
		token.Value = readMap(r)
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
