package ringbuffer

import goutils "github.com/felipe-bonato/go-utils"

type RingBuffer[T any] struct {
	values []T
	len    int
	write  int // Index into ´values´ the first item
	read   int // Index into ´values´ the last item
}

func New[T any](size int) RingBuffer[T] {
	if size <= 0 {
		panic("ringbuffer: size must be > ")
	}
	return RingBuffer[T]{
		values: make([]T, size),
		len:    0,
		write:  0,
		read:   0,
	}
}

// Inserts ´value´ in the front of the ring buffer.
// If the buffer is full, overwrites the oldest  element,
// retuning it and setting `evicetedOk` to true, else returns `_, false`
func (rb *RingBuffer[T]) Push(value T) (evicted T, evictedOk bool) {
	if rb == nil || cap(rb.values) == 0 {
		panic("ringbuffer: uninitialized `Push` use")
	}

	if rb.len > 0 && rb.read == rb.write {
		rb.read = (rb.read + 1) % cap(rb.values)
		evicted = rb.values[rb.write]
		evictedOk = true
	} else {
		rb.len += 1
	}

	rb.values[rb.write] = value

	rb.write = (rb.write + 1) % cap(rb.values)

	return evicted, evictedOk
}

func (rb *RingBuffer[T]) Pop() (value T, ok bool) {
	if rb == nil || cap(rb.values) == 0 {
		panic("ringbuffer: uninitialized `Pop()` use")
	}

	if rb.len == 0 {
		return value, false
	}

	value = rb.values[rb.read]

	// Zero it out, so if it is a reference, the GC can dealocate it.
	rb.values[rb.read] = goutils.Empty[T]()

	rb.read = (rb.read + 1) % cap(rb.values)
	rb.len--

	return value, true
}

func (rb *RingBuffer[T]) Len() int {
	if rb == nil {
		return 0
	}

	return rb.len
}
