package model

import "time"

type ArrayEncoder interface {
	PutArray(int) (ArrayEncoder, error)
	PutDate(time.Time) error
	PutInt(int) error
	PutMap(int) (MapEncoder, error)
	PutObject() (ObjectEncoder, error)
	PutRef(Ref) error
	PutString(string) error
}

type MapEncoder interface {
	PutArray(int) (ArrayEncoder, error)
	PutDate(time.Time) error
	PutInt(int) error
	PutMap(int) (MapEncoder, error)
	PutObject() (ObjectEncoder, error)
	PutRef(Ref) error
	PutString(string) error
}

type ObjectEncoder interface {
	PutArray(string, int) (ArrayEncoder, error)
	PutDate(string, time.Time) error
	PutInt(string, int) error
	PutMap(string, int) (MapEncoder, error)
	PutObject(string) (ObjectEncoder, error)
	PutRef(string, Ref) error
	PutString(string, string) error
}
