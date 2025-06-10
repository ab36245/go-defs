package codec

import "time"

type ArrayTarget interface {
	PutArray(int, func(ArrayTarget))
	PutDate(time.Time)
	PutInt(int)
	PutObject(func(ObjectTarget))
	PutString(string)
}

type ObjectTarget interface {
	PutArray(string, int, func(ArrayTarget))
	PutDate(string, time.Time)
	PutInt(string, int)
	PutObject(string, func(ObjectTarget))
	PutString(string, string)
}
