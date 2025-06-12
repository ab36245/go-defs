package defs

type Codec[T any] struct {
	Decode func(ObjectDecoder) (T, error)
	Encode func(ObjectEncoder, T)
}
