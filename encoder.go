package model

import "time"

type ArrayEncoder interface {
	PutArray(int, func(ArrayEncoder))
	PutDate(time.Time)
	PutInt(int)
	PutObject(func(ObjectEncoder))
	PutRef(Ref)
	PutString(string)
}

type ObjectEncoder interface {
	PutArray(string, int, func(ArrayEncoder))
	PutDate(string, time.Time)
	PutInt(string, int)
	PutObject(string, func(ObjectEncoder))
	PutRef(string, Ref)
	PutString(string, string)
}
