package model

type Option[T any] struct {
	isSet bool
	value T
}
