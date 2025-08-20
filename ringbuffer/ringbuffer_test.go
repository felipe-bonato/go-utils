package ringbuffer_test

import (
	"testing"

	"github.com/felipe-bonato/goutils/ringbuffer"
)

func TestRingBuffer(t *testing.T) {
	rb := ringbuffer.New[int](3)

	expectLen := func(expectedLen int) {
		if rb.Len() != expectedLen {
			t.Errorf("expected len %d got %d", expectedLen, rb.Len())
		}
	}

	pushLen := func(value int, expectedLen int) {
		rb.Push(value)
		expectLen(expectedLen)
	}

	pushRemoved := func(value int, expectedLen int, expectedRemoved int) {
		item, removed := rb.Push(value)
		expectLen(expectedLen)

		if !removed {
			t.Errorf("expected %d to be removed", expectedRemoved)
		}
		if item != expectedRemoved {
			t.Errorf("expected %d to be removed, but got %d", expectedRemoved, item)
		}
	}

	pop := func(expectedValue int, expectedOk bool, expectedLen int) {
		v, ok := rb.Pop()
		if ok != expectedOk || v != expectedValue {
			t.Errorf("Expected to %d %t, but got %d %t", expectedValue, expectedOk, v, ok)
		}
		expectLen(expectedLen)
	}

	pop(0, false, 0)

	pushLen(1, 1)
	pop(1, true, 0)
	pushLen(1, 1)

	pushLen(2, 2)
	pushLen(3, 3)

	pushRemoved(4, 3, 1)
	pushRemoved(5, 3, 2)
	pushRemoved(6, 3, 3)

	pushRemoved(7, 3, 4)

	pop(5, true, 2)
	pop(6, true, 1)
	pop(7, true, 0)
	pop(0, false, 0)

	// t.Logf("internal: %+v", rb)
}
