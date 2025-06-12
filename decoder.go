package defs

import "time"

type ArrayDecoder interface {
	GetArray() (ArrayDecoder, error)
	GetDate() (time.Time, error)
	GetInt() (int, error)
	GetString() (string, error)
	GetObject() (ObjectDecoder, error)
	Length() int
}

type ObjectDecoder interface {
	GetArray(string) (ArrayDecoder, error)
	GetDate(string) (time.Time, error)
	GetInt(string) (int, error)
	GetString(string) (string, error)
	GetObject(string) (ObjectDecoder, error)
}
