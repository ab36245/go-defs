package model

import "time"

type ArrayEncoder interface {
	PutArray(int, func(ArrayEncoder))
	PutDate(time.Time)
	PutInt(int)
	PutMap(int, func(MapEncoder))
	PutObject(func(ObjectEncoder))
	PutRef(Ref)
	PutString(string)
}

type MapEncoder interface {
	PutArray(int, func(ArrayEncoder))
	PutDate(time.Time)
	PutInt(int)
	PutKey(string)
	PutMap(int, func(MapEncoder))
	PutObject(func(ObjectEncoder))
	PutRef(Ref)
	PutString(string)
}

type ObjectEncoder interface {
	PutArray(string, int, func(ArrayEncoder))
	PutDate(string, time.Time)
	PutInt(string, int)
	PutMap(string, int, func(MapEncoder))
	PutObject(string, func(ObjectEncoder))
	PutRef(string, Ref)
	PutString(string, string)
}
