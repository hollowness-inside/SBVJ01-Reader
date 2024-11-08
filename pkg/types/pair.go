package types

import "fmt"

type SBVJPair struct {
	Key   string
	Value SBVJObject
}

func (p *SBVJPair) String() string {
	return fmt.Sprintf(`"%s": %v`, p.Key, p.Value)
}
