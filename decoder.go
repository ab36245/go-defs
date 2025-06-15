package model

import "time"

type ArrayDecoder interface {
	GetArray() (ArrayDecoder, error)
	GetDate() (time.Time, error)
	GetInt() (int, error)
	GetMap() (MapDecoder, error)
	GetObject() (ObjectDecoder, error)
	GetRef() (Ref, error)
	GetString() (string, error)
	Length() int
}

type MapDecoder interface {
	GetArray() (ArrayDecoder, error)
	GetDate() (time.Time, error)
	GetInt() (int, error)
	GetMap() (MapDecoder, error)
	GetObject() (ObjectDecoder, error)
	GetRef() (Ref, error)
	GetString() (string, error)
	Length() int
}

type ObjectDecoder interface {
	GetArray(string) (ArrayDecoder, error)
	GetDate(string) (time.Time, error)
	GetInt(string) (int, error)
	GetMap(string) (MapDecoder, error)
	GetObject(string) (ObjectDecoder, error)
	GetRef(string) (Ref, error)
	GetString(string) (string, error)
}
