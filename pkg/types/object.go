package types

import (
	"fmt"
	"strings"
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
		return formatList(o.Value.(SBVJList))
	case MAP:
		return formatMap(o.Value.(SBVJMap))
	default:
		return ""
	}
}

func formatList(l SBVJList) string {
	strs := make([]string, len(l))

	for i, v := range l {
		strs[i] = v.String()
	}

	return fmt.Sprintf("[%s]", strings.Join(strs, ", "))
}

func formatMap(m SBVJMap) string {
	elements := make([]string, 0, len(m))

	for key, value := range m {
		elements = append(elements, fmt.Sprintf(`"%s": %s`, key, value.String()))
	}

	return fmt.Sprintf("{%s}", strings.Join(elements, ", "))
}
