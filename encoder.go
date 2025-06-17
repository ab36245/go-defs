package model

import "time"

type ArrayEncoder interface {
	PutArray(int) (ArrayEncoder, error)
	PutBool(bool) error
	PutFloat(float64) error
	PutInt(int) error
	PutMap(int) (MapEncoder, error)
	PutObject() (ObjectEncoder, error)
	PutRef(Ref) error
	PutString(string) error
	PutTime(time.Time) error
}

type MapEncoder interface {
	PutArray(int) (ArrayEncoder, error)
	PutBool(bool) error
	PutFloat(float64) error
	PutInt(int) error
	PutMap(int) (MapEncoder, error)
	PutObject() (ObjectEncoder, error)
	PutRef(Ref) error
	PutString(string) error
	PutTime(time.Time) error
}

type ObjectEncoder interface {
	PutArray(string, int) (ArrayEncoder, error)
	PutBool(string, bool) error
	PutFloat(string, float64) error
	PutInt(string, int) error
	PutMap(string, int) (MapEncoder, error)
	PutObject(string) (ObjectEncoder, error)
	PutRef(string, Ref) error
	PutString(string, string) error
	PutTime(string, time.Time) error
}
