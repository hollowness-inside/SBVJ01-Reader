package sbvj01

import "fmt"

type SBVJ01Pair struct {
	Key   string
	Value *SBVJ01Token
}

func (p *SBVJ01Pair) String() string {
	return fmt.Sprintf("\"%s\": %v", p.Key, p.Value)
}
