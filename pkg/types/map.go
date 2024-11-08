package types

import (
	"fmt"
	"strings"
)

type SBVJMap struct {
	Items map[string]SBVJObject
}

func (m *SBVJMap) Get(key string) *SBVJObject {
	v := m.Items[key]
	return &v
}

func (m *SBVJMap) Set(key string, value SBVJObject) {
	m.Items[key] = value
}

func (m *SBVJMap) String() string {
	elements := make([]string, 0, len(m.Items))

	for i, v := range m.Items {
		elements = append(elements, fmt.Sprintf(`"%s": %s`, i, v.String()))
	}

	return fmt.Sprintf("{%s}", strings.Join(elements, ", "))
}
