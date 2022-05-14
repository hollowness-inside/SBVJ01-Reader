package sbvj01

import (
	"bufio"
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
}
