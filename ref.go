package model

import (
	"fmt"
)

type Ref string

func (r Ref) String() string {
	return fmt.Sprintf("%T(%s)", r, string(r))
}
