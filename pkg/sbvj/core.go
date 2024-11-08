package sbvj

import (
	"github.com/hollowness-inside/SBVJ01-Reader/pkg/types"
)

type SBVJ struct {
	Name      string
	Versioned bool
	Version   int32
	Content   types.SBVJObject
}
