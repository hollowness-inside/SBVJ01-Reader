package sbvj01

import (
	"fmt"
)

type SBVJ01Token struct {
	Type  byte
	Value any
}

func (t SBVJ01Token) String() string {
	switch t.Type {
	case NIL:
		return "nil"
	case DOUBLE:
		return fmt.Sprintf("%f", t.Value.(float64))
	case BOOLEAN:
		return fmt.Sprintf("%t", t.Value.(bool))
	case VARINT:
		return fmt.Sprintf("%d", t.Value.(int))
	case STRING:
		return fmt.Sprintf("\"%s\"", t.Value.(string))
	case LIST:
		return fmt.Sprintf("%v", t.Value)
	case MAP:
		return t.Value.(SBVJ01Map).String()
	}

	return ""
}
