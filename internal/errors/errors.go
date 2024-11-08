package errors

import (
	"fmt"

	"github.com/hollowness-inside/SBVJ01-Reader/pkg/types"
)

type ErrMagic struct {
	magic []byte
}

func NewErrMagic(received []byte) *ErrMagic {
	return &ErrMagic{
		magic: received,
	}
}

type ErrObjectType struct {
	t types.SBVJType
}

func NewErrObjectType(t types.SBVJType) *ErrObjectType {
	return &ErrObjectType{
		t: t,
	}
}

func (e *ErrMagic) Error() string {
	return fmt.Sprintf("got wrong magic - expected SBVJ01, received %s", string(e.magic))
}

func (e *ErrObjectType) Error() string {
	return fmt.Sprintf("unknown object type (%s)", string(e.t))
}
