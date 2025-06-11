package codec

import "time"

type ArrayEncoder interface {
	PutArray(int, func(ArrayEncoder))
	PutDate(time.Time)
	PutInt(int)
	PutObject(func(ObjectEncoder))
	PutString(string)
}

type ObjectEncoder interface {
	PutArray(string, int, func(ArrayEncoder))
	PutDate(string, time.Time)
	PutInt(string, int)
	PutObject(string, func(ObjectEncoder))
	PutString(string, string)
}
