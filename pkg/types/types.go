package types

type SBVJType byte

const (
	NIL     SBVJType = 0x01
	DOUBLE  SBVJType = 0x02
	BOOLEAN SBVJType = 0x03
	VARINT  SBVJType = 0x04
	STRING  SBVJType = 0x05
	LIST    SBVJType = 0x06
	MAP     SBVJType = 0x07
)
