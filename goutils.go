package goutils

import (
	"errors"
	"fmt"
	"os"
	"golang.org/x/exp/constraints"
)

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

// Clamps a  `value` between `min` and `max`.
func Clamp[T Number](value, min, max T) T {
	if value < min {
		value = min
	} else if value > max {
		value = max
	}
	return value
}

// Gets deref of `value` or `defaultValue` if nil
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

// Returns true if exists, false if it doesn't, and false and error if unknowable.
func FileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		// It's only guaranteed that the file doesn't exists if the return is ErrNotExist.
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}

		// Schrodinger file: it may or may not exist.
		// See for more info: https://stackoverflow.com/a/12518877
		return false, fmt.Errorf("cannot know if file exists: %w", err)
	}

	return true, nil
}
