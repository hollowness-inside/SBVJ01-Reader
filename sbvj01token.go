package sbvj01

import "fmt"

type SBVJ01Token struct {
	Type  byte
	Value any
}

func (t SBVJ01Token) String() string {
	if t.Type == NIL {
		return "NIL"
	}

	var typeStr string
	switch t.Type {
	case DOUBLE:
		typeStr = "DOUBLE"
	case BOOLEAN:
		typeStr = "BOOLEAN"
	case VARINT:
		typeStr = "VARINT"
	case STRING:
		typeStr = "STRING"
	case LIST:
		typeStr = "LIST"
	case MAP:
		typeStr = "MAP"
	}
	return fmt.Sprintf("%s(%v)", typeStr, t.Value)
}
