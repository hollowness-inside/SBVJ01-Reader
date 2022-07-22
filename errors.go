package sbvj

import "fmt"

type ErrMagic struct {
	magic []byte
}

type ErrObjectType struct {
	t SBVJType
}

func (e *ErrMagic) Error() string {
	return fmt.Sprintf("got wrong magic - expected SBVJ01, received %s", string(e.magic))
}

func (e *ErrObjectType) Error() string {
	return fmt.Sprintf("unknown object type (%s)", string(e.t))
}
