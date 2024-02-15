package ringbuf

// Buffer provides ring buffer.
type Buffer[T any] struct {
	a  []T
	rx int
	n  int
}

// New creates a ring buffer with specified capacity.
func New[T any](capacity int) *Buffer[T] {
	if capacity <= 0 {
		panic("capacity must be large than 0")
	}
	return &Buffer[T]{
		a: make([]T, capacity),
	}
}

// Enqueue puts a value to last of the buffer.
func (b *Buffer[T]) Enqueue(v T) {
	wx := b.rx + b.n
	if wx >= len(b.a) {
		wx -= len(b.a)
	}
	b.a[wx] = v
	if b.n < len(b.a) {
		b.n++
		return
	}
	b.rxInc()
}

// Dequeue retrieves a value.
func (b *Buffer[T]) Dequeue() (T, bool) {
	var zero T
	if b.n <= 0 {
		return zero, false
	}
	v := b.a[b.rx]
	b.a[b.rx] = zero
	b.rxInc()
	b.n--
	return v, true
}

// Clear remove all values.
func (b *Buffer[T]) Clear() {
	var zero T
	for i := range b.a {
		b.a[i] = zero
	}
	b.rx = 0
	b.n = 0
}

// Empty checks the buffer is empty or not.
func (b *Buffer[T]) Empty() bool {
	return b.n == 0
}

// rxInc increments rx.
func (b *Buffer[T]) rxInc() int {
	b.rx++
	if b.rx == len(b.a) {
		b.rx = 0
	}
	return b.rx
}

// Cap returns capacity of the ringbuf.
func (b *Buffer[T]) Cap() int {
	return len(b.a)
}

// Len returns number of valid items in the ringbuf.
func (b *Buffer[T]) Len() int {
	return b.n
}

// Peek peeks a n'th value in the ringbuf without removing.
func (b *Buffer[T]) Peek(n int) T {
	var zero T
	if n < 0 || n >= b.n {
		return zero
	}
	rx := b.rx + n
	if rx >= len(b.a) {
		rx -= len(b.a)
	}
	return b.a[rx]
}
