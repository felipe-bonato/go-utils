package ringbuffer

type RingBuffer[T any] struct {
	values []T
	len    int
	write  int // Index into ´values´ the first item
	read   int // Index into ´values´ the last item
}

func New[T any](size int) RingBuffer[T] {
	return RingBuffer[T]{
		values: make([]T, size, size),
		len:    0,
		write:  0,
		read:   0,
	}
}

// Inserts ´value´ in the front of the ring buffer,
// returning ´T, true´ if an item was removed to insert it.
func (rb *RingBuffer[T]) PushBack(value T) (item T, removed bool) {
	if rb.len > 0 && rb.read == rb.write {
		rb.read = (rb.read + 1) % cap(rb.values)
		item = rb.values[rb.write]
		removed = true
	} else {
		rb.len += 1
	}

	rb.values[rb.write] = value

	rb.write = (rb.write + 1) % cap(rb.values)

	return item, removed
}

func (rb *RingBuffer[T]) PopFront() (value T, ok bool) {
	if rb.len == 0 {
		return value, false
	}

	value = rb.values[rb.read]

	rb.read = (rb.read + 1) % cap(rb.values)
	rb.len--

	return value, true
}

func (rb *RingBuffer[T]) Len() int {
	return rb.len
}
