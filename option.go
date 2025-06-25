package model

func EmptyOption[T any]() Option[T] {
	return Option[T]{}
}

func NewOption[T any](value T) Option[T] {
	return Option[T]{
		isSet: true,
		value: value,
	}
}

type Option[T any] struct {
	isSet bool
	value T
}

func (o Option[T]) IsSet() bool {
	return o.isSet
}

func (o Option[T]) Value() T {
	return o.value
}
