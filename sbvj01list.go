package sbvj01

import (
	"fmt"
	"strings"
)

type SBVJ01List struct {
	Items []SBVJ01Token
}

func (l *SBVJ01List) Get(n int) *SBVJ01Token {
	return &l.Items[n]
}

func (l SBVJ01List) String() string {
	strs := make([]string, len(l.Items))
	for i, v := range l.Items {
		strs[i] = v.String()
	}

	return fmt.Sprintf("[%s]", strings.Join(strs, ", "))
}
