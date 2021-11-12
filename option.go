package option

import "errors"

type Option[T any] interface {
	isSome() bool
	Get() T
}

type option[T any] struct {
	value T
	some  bool
}

func (o option[T]) isSome() bool {
	return o.some
}

func (o option[T]) Get() T {
	if o.isSome() {
		return o.value
	}
	panic(errors.New("Option.Get() called on None"))
}

func Some[T any](v T) Option[T] {
	return option[T]{value: v, some: true}
}

func None[T any]() Option[T] {
	return option[T]{some: false}
}

func Map[T any, R any](m func(T) R, o Option[T]) Option[R] {
	if o.isSome() {
		return Some(m(o.Get()))
	} else {
		// better would be but not working per now
		// o.(Option[R])
		return None[R]()
	}
}

// panic on none
func Get[T any](o Option[T]) T {
	return o.Get()
}

func Or[T any](or T, o Option[T]) T {
	if o.isSome() {
		return o.Get()
	} else {
		return or
	}
}

func OrElse[T any](or func() T, o Option[T]) T {
	if o.isSome() {
		return o.Get()
	} else {
		return or()
	}
}

func IsSome[T any](o Option[T]) bool {
	return o.isSome()
}

func IsNone[T any](o Option[T]) bool {
	return !o.isSome()
}
