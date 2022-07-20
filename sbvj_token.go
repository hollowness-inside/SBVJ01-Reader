package sbvj

import (
	"fmt"
)

type SBVJToken struct {
	Type  SBVJType
	Value any
}

func (t SBVJToken) String() string {
	switch t.Type {
	case NIL:
		return "null"
	case DOUBLE:
		return fmt.Sprintf("%f", t.Value.(float64))
	case BOOLEAN:
		return fmt.Sprintf("%t", t.Value.(bool))
	case VARINT:
		return fmt.Sprintf("%d", t.Value.(int64))
	case STRING:
		return fmt.Sprintf(`"%s"`, t.Value.(string))
	case LIST:
		return t.Value.(SBVJList).String()
	case MAP:
		return t.Value.(SBVJMap).String()
	default:
		return ""
	}
}
