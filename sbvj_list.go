package sbvj

import (
	"fmt"
	"strings"
)

type SBVJList struct {
	Items []SBVJToken
}

func (l *SBVJList) Get(n int) *SBVJToken {
	return &l.Items[n]
}

func (l SBVJList) String() string {
	strs := make([]string, len(l.Items))
	for i, v := range l.Items {
		strs[i] = v.String()
	}

	return fmt.Sprintf("[%s]", strings.Join(strs, ", "))
}
