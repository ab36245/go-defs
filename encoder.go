package model

import "time"

type ArrayHandler func(ArrayEncoder) error
type MapHandler func(MapEncoder) error
type ObjectHandler func(ObjectEncoder) error

type ArrayEncoder interface {
	PutArray(int, ArrayHandler) error
	PutDate(time.Time) error
	PutInt(int) error
	PutMap(int, MapHandler) error
	PutObject(ObjectHandler) error
	PutRef(Ref) error
	PutString(string) error
}

type MapEncoder interface {
	PutArray(int, ArrayHandler) error
	PutDate(time.Time) error
	PutInt(int) error
	PutMap(int, MapHandler) error
	PutObject(ObjectHandler) error
	PutRef(Ref) error
	PutString(string) error
}

type ObjectEncoder interface {
	PutArray(string, int, ArrayHandler) error
	PutDate(string, time.Time) error
	PutInt(string, int) error
	PutMap(string, int, MapHandler) error
	PutObject(string, ObjectHandler) error
	PutRef(string, Ref) error
	PutString(string, string) error
}
