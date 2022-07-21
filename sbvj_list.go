package sbvj

import (
	"fmt"
	"strings"
)

type SBVJList struct {
	Items []SBVJObject
}

func (l *SBVJList) Get(n int) *SBVJObject {
	return &l.Items[n]
}

func (l *SBVJList) String() string {
	strs := make([]string, len(l.Items))
	for i, v := range l.Items {
		strs[i] = v.String()
	}

	return fmt.Sprintf("[%s]", strings.Join(strs, ", "))
}
