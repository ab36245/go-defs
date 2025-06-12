package model

import "time"

type ArrayDecoder interface {
	GetArray() (ArrayDecoder, error)
	GetDate() (time.Time, error)
	GetInt() (int, error)
	GetObject() (ObjectDecoder, error)
	GetRef() (Ref, error)
	GetString() (string, error)
	Length() int
}

type ObjectDecoder interface {
	GetArray(string) (ArrayDecoder, error)
	GetDate(string) (time.Time, error)
	GetInt(string) (int, error)
	GetObject(string) (ObjectDecoder, error)
	GetRef(string) (Ref, error)
	GetString(string) (string, error)
}
