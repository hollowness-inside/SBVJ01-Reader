package sbvj01

import (
	"fmt"
	"strings"
)

type SBVJ01Token struct {
	Type  byte
	Value any
}

func (t SBVJ01Token) String() string {
	if t.Type == NIL {
		return "NIL"
	} else if t.Type == DOUBLE {
		return fmt.Sprintf("%f", t.Value)
	} else if t.Type == MAP {
		mapValue := t.Value.([]*SBVJ01Pair)
		elements := make([]string, len(mapValue))

		for i, v := range mapValue {
			elements[i] = v.String()
		}

		return "{" + strings.Join(elements, ", ") + "}"
	}

	var typeStr string
	switch t.Type {
	case BOOLEAN:
		typeStr = "BOOLEAN"
	case VARINT:
		typeStr = "VARINT"
	case STRING:
		typeStr = "STRING"
	case LIST:
		typeStr = "LIST"
	}

	return fmt.Sprintf("%s(%v)", typeStr, t.Value)
}
