package sbvj

import (
	"fmt"
)

type SBVJObject struct {
	Type  SBVJType
	Value any
}

func (o SBVJObject) String() string {
	switch o.Type {
	case NIL:
		return "null"
	case DOUBLE:
		return fmt.Sprintf("%f", o.Value.(float64))
	case BOOLEAN:
		return fmt.Sprintf("%t", o.Value.(bool))
	case VARINT:
		return fmt.Sprintf("%d", o.Value.(int64))
	case STRING:
		return fmt.Sprintf(`"%s"`, o.Value.(string))
	case LIST:
		l := o.Value.(SBVJList)
		return l.String()
	case MAP:
		m := o.Value.(SBVJMap)
		return m.String()
	default:
		return ""
	}
}
