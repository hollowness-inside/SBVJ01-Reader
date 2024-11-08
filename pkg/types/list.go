package types

import (
	"fmt"
	"strings"
)

type SBVJList struct {
	Items []SBVJObject
}

func (l *SBVJList) At(n int) *SBVJObject {
	if n < 0 || n >= len(l.Items) {
		return nil
	}

	return &l.Items[n]
}

func (l *SBVJList) Append(item SBVJObject) {
	l.Items = append(l.Items, item)
}

func (l *SBVJList) Len() int {
	return len(l.Items)
}

func (l *SBVJList) Cap() int {
	return cap(l.Items)
}

func (l SBVJList) String() string {
	strs := make([]string, len(l.Items))
	for i, v := range l.Items {
		strs[i] = v.String()
	}

	return fmt.Sprintf("[%s]", strings.Join(strs, ", "))
}
