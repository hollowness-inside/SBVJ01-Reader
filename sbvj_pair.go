package sbvj

import "fmt"

type SBVJPair struct {
	Key   string
	Value SBVJToken
}

func (p SBVJPair) String() string {
	return fmt.Sprintf(`"%s": %v`, p.Key, p.Value)
}
