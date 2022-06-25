package sbvj01

import "strings"

type SBVJ01Map struct {
	Items []SBVJ01Pair
}

func (m SBVJ01Map) String() string {
	elements := make([]string, len(m.Items))

	for i, v := range m.Items {
		elements[i] = v.String()
	}

	return "{" + strings.Join(elements, ", ") + "}"
}

func (m *SBVJ01Map) Get(key string) *SBVJ01Token {
	for _, it := range m.Items {
		if it.Key == key {
			return &it.Value
		}
	}

	return nil
}
