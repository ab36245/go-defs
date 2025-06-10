package codec

type Codec[T any] struct {
	Decode func(ObjectSource) (T, error)
	Encode func(ObjectTarget, T)
}
