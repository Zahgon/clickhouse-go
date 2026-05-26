package circular

import "iter"

// Queue is a bounded FIFO queue implemented using a circular array.
// It uses head and tail pointers to avoid slice re-allocations.
// When full, new elements are rejected rather than overwriting old ones.
type Queue[T any] struct {
	data []T
	head int // index of the first element
	tail int // index where the next element will be inserted
	len  int // number of elements in the queue
}

// New creates a new circular queue with the given capacity.
func New[T any](capacity int) *Queue[T] { _ = "STUB: not implemented"; return nil }

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int {
	_ = "STUB: not implemented"

	// Cap returns the capacity of the queue.
	return 0
}

func (q *Queue[T]) Cap() int {
	_ = "STUB: not implemented"

	// IsFull returns true if the queue is at capacity.
	return 0
}

func (q *Queue[T]) IsFull() bool { _ = "STUB: not implemented"; return false }

// IsEmpty returns true if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	_ = "STUB: not implemented"

	// Push adds an element to the tail of the queue.
	// Returns false if the queue is full.
	return false
}

func (q *Queue[T]) Push(value T) bool { _ = "STUB: not implemented"; return false }

// Pull removes and returns an element from the head of the queue.
// Returns the zero value and false if the queue is empty.
func (q *Queue[T]) Pull() (value T, ok bool) { _ = "STUB: not implemented"; return *new(T), false }

// all returns an iterator over all elements in the queue in FIFO order.
// The iterator yields (index, value) pairs where index is 0-based from the head.
func (q *Queue[T]) all() iter.Seq2[int, T] { _ = "STUB: not implemented"; return nil }

// DeleteFunc removes elements from the queue based on a predicate function.
// Returns an iterator over the removed elements.
// Elements for which shouldRemove returns true are removed from the queue.
func (q *Queue[T]) DeleteFunc(shouldRemove func(T) bool) (removed iter.Seq[T]) {
	_ = "STUB: not implemented"
	return nil
}

// Keep this element - move it to newTail if needed

// Remove this element

// Try to yield the removed value if we haven't stopped

// Clear removes all elements from the queue.
// Returns an iterator over the removed elements.
func (q *Queue[T]) Clear() iter.Seq[T] { _ = "STUB: not implemented"; return nil }

// next returns the next index in the circular queue.
func (q *Queue[T]) next(index int) int { _ = "STUB: not implemented"; return 0 }
