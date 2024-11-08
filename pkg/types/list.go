package types

import (
	"fmt"
	"strings"
)

type SBVJList []SBVJObject

func (l SBVJList) String() string {
	strs := make([]string, len(l))
	for i, v := range l {
		strs[i] = v.String()
	}

	return fmt.Sprintf("[%s]", strings.Join(strs, ", "))
}
