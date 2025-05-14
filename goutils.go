package goutils

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Same as `&v`.
// Direct pointer return, useful for untyped types, because you cannot take reference to then.
func Ptr[T any](v T) *T {
	return &v
}

// Returns the empty representation of a value.
func Empty[T any]() T {
	var x T
	return x
}

// For prototyping purposes.
func Try(err error) {
	if err != nil {
		panic(err)
	}
}

func Clamp[T Number](value, min, max T) T {
	if value < min {
		value = min
	} else if value > max {
		value = max
	}
	return value
}

func Default[T any](value *T, defaultValue T) T {
	if value == nil {
		return defaultValue
	}
	return *value
}

// Ternary operator.
// Same as `condition ? ifTrue : ifFalse`.
func Ternary[T any](condition bool, ifTrue T, ifFalse T) T {
	if condition {
		return ifTrue
	} else {
		return ifFalse
	}
}

// Short circuit "Or".
// Same as `first || second`.
// Uses the empty value of first for comparison.
func Or[T comparable](first T, second T) T {
	if first == Empty[T]() {
		return second
	} else {
		return first
	}
}

// Short circuit "And".
// Same as `first && second`.
// Uses the empty value of first for comparison.
func And[T comparable](first T, second T) T {
	if first == Empty[T]() {
		return first
	} else {
		return second
	}
}
