package types

import (
	"fmt"
	"strings"
)

type SBVJMap map[string]SBVJObject

func (m SBVJMap) String() string {
	elements := make([]string, 0, len(m))

	for i, v := range m {
		elements = append(elements, fmt.Sprintf(`"%s": %s`, i, v.String()))
	}

	return fmt.Sprintf("{%s}", strings.Join(elements, ", "))
}
