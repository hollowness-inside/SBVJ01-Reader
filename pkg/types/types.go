package types

type SBVJType byte

type SBVJList = []SBVJObject
type SBVJMap = map[string]SBVJObject

const (
	_ SBVJType = iota
	NIL
	DOUBLE
	BOOLEAN
	VARINT
	STRING
	LIST
	MAP
)
