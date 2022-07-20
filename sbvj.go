package sbvj

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type SBVJType byte

const (
	NIL     SBVJType = 0x01
	DOUBLE  SBVJType = 0x02
	BOOLEAN SBVJType = 0x03
	VARINT  SBVJType = 0x04
	STRING  SBVJType = 0x05
	LIST    SBVJType = 0x06
	MAP     SBVJType = 0x07
)

type SBVJ struct {
	Name      string
	Versioned bool
	Version   int32
	Value     SBVJToken
}

func ReadBytes(buf []byte) SBVJ {
	buffer := bytes.NewBuffer(buf)
	return Read(buffer)
}

func ReadSBVJFile(path string) SBVJ {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return Read(file)
}

func Read(r io.Reader) SBVJ {
	reader := bufio.NewReader(r)

	magic := make([]byte, 6)
	_, err := io.ReadFull(reader, magic)
	if err != nil {
		panic(err)
	}

	if string(magic) != "SBVJ01" {
		panic("Wrong magic")
	}

	sbvj := SBVJ{}
	sbvj.Name = readString(reader)

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
	_, err := io.ReadFull(r, bytes)
	if err != nil {
		panic(err)
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

func readToken(r *bufio.Reader) SBVJToken {
	token := SBVJToken{}

	tp := readByte(r)
	token.Type = SBVJType(tp)

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

func readList(r *bufio.Reader) SBVJList {
	sbvjList := SBVJList{}

	size := readVarint(r)
	sbvjList.Items = make([]SBVJToken, size)

	var i int64
	for i = 0; i < size; i++ {
		token := readToken(r)
		sbvjList.Items[i] = token
	}

	return sbvjList
}

func readMap(r *bufio.Reader) SBVJMap {
	sbvjmap := SBVJMap{}

	size := readVarint(r)
	sbvjmap.Items = make([]SBVJPair, size)

	var i int64
	for i = 0; i < size; i++ {
		key := readString(r)
		value := readToken(r)
		sbvjmap.Items[i] = SBVJPair{key, value}
	}

	return sbvjmap
}
