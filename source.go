package codec

import "time"

type ArraySource interface {
	GetArray() (ArraySource, error)
	GetDate() (time.Time, error)
	GetInt() (int, error)
	GetString() (string, error)
	GetObject() (ObjectSource, error)
	Length() int
}

type ObjectSource interface {
	GetArray(string) (ArraySource, error)
	GetDate(string) (time.Time, error)
	GetInt(string) (int, error)
	GetString(string) (string, error)
	GetObject(string) (ObjectSource, error)
}
