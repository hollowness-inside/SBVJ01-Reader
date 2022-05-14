package sbvj01

import "strings"

type SBVJ01List struct {
	Items []SBVJ01Token
}

func (l SBVJ01List) String() string {
	strs := make([]string, len(l.Items))
	for i, v := range l.Items {
		strs[i] = v.String()
	}

	return "[" + strings.Join(strs, ", ") + "]"
}

func (l *SBVJ01List) Get(n int) *SBVJ01Token {
	return &l.Items[n]
}

func NewSBVJ01List(size int) *SBVJ01List {
	list := new(SBVJ01List)
	list.Items = make([]SBVJ01Token, size)
	return list
}
