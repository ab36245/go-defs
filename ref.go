package model

import "fmt"

type Ref string

func (r Ref) String() string {
	return fmt.Sprintf("Ref(%s)", r)
}
