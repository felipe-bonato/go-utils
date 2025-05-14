package option

import (
	"encoding/json"
)

type Option[T any] struct {
	val T
	ok  bool
}

func Some[T any](val T) Option[T] {
	return Option[T]{
		val: val,
		ok:  true,
	}
}

func None[T any]() Option[T] {
	var empty T
	return Option[T]{
		val: empty,
		ok:  false,
	}
}

func (opt Option[T]) Val() (T, bool) {
	return opt.val, opt.ok
}

func (opt Option[T]) Or(val T) T {
	if !opt.ok {
		return val
	}

	return opt.val
}

func Flatten[T any](opt Option[Option[T]]) Option[T] {
	if !opt.ok {
		return None[T]()
	}
	return opt.val
}

// Use this with go's 1.24 `omitzero` tag.
func (opt *Option[T]) UnmarshalJSON(d []byte) error {
	if err := json.Unmarshal(d, &opt.val); err != nil {
		return err
	}

	opt.ok = true

	return nil
}

// Use this with go's 1.24 `omitzero` tag.
func (opt Option[T]) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(opt.val)
	if err != nil {
		return nil, err
	}

	return data, nil
}
