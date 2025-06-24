package model

import (
	"fmt"
	"time"
)

type Ref string

func (r Ref) String() string {
	return fmt.Sprintf("%T(%s)", r, string(r))
}

type Date time.Time

func (d Date) String() string {
	return fmt.Sprintf("%T(%s)", d, time.Time(d).Format(time.DateOnly))
}

type Time time.Time

func (t Time) String() string {
	return fmt.Sprintf("%T(%s)", t, time.Time(t).Format(time.TimeOnly))
}
