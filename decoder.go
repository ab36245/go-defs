package model

import "time"

type ArrayDecoder interface {
	GetArray() (ArrayDecoder, error)
	GetBool() (bool, error)
	GetBytes() ([]byte, error)
	GetFloat() (float64, error)
	GetInt() (int, error)
	GetMap() (MapDecoder, error)
	GetObject() (ObjectDecoder, error)
	GetRef() (Ref, error)
	GetString() (string, error)
	GetTime() (time.Time, error)
	Length() int
}

type MapDecoder interface {
	GetArray() (ArrayDecoder, error)
	GetBool() (bool, error)
	GetBytes() ([]byte, error)
	GetFloat() (float64, error)
	GetInt() (int, error)
	GetMap() (MapDecoder, error)
	GetObject() (ObjectDecoder, error)
	GetRef() (Ref, error)
	GetString() (string, error)
	GetTime() (time.Time, error)
	Length() int
}

type ObjectDecoder interface {
	GetArray(string) (ArrayDecoder, error)
	GetBool(string) (bool, error)
	GetBytes(string) ([]byte, error)
	GetFloat(string) (float64, error)
	GetInt(string) (int, error)
	GetMap(string) (MapDecoder, error)
	GetObject(string) (ObjectDecoder, error)
	GetRef(string) (Ref, error)
	GetString(string) (string, error)
	GetTime(string) (time.Time, error)
}
