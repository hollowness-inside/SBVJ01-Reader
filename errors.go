package sbvj

import "fmt"

type ErrWrongMagic struct {
	magic []byte
}

type ErrUnknownObjectType struct {
	t SBVJType
}

func (e *ErrWrongMagic) Error() string {
	return fmt.Sprintf("got wrong magic - expected SBVJ01, received %s", string(e.magic))
}

func (e *ErrUnknownObjectType) Error() string {
	return fmt.Sprintf("unknown object type (%s)", string(e.t))
}
