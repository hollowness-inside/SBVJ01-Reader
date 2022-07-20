package sbvj

import (
	"fmt"
	"strings"
)

type SBVJMap struct {
	Items []SBVJPair
}

func (m *SBVJMap) Get(key string) *SBVJToken {
	for _, it := range m.Items {
		if it.Key == key {
			return &it.Value
		}
	}

	return nil
}

func (m SBVJMap) String() string {
	elements := make([]string, len(m.Items))

	for i, v := range m.Items {
		elements[i] = v.String()
	}

	return fmt.Sprintf("{%s}", strings.Join(elements, ", "))
}
