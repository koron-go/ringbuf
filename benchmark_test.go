package ringbuf

import (
	"testing"
)

func BenchmarkEnqueueDequeue(b *testing.B) {
	rb := New[int](2 * 1024 * 1024)
	b.ResetTimer()
	for n := max(1, b.N/2); n > 0; n -= 1000 {
		m := min(n, 1000)
		for i := 0; i < m; i++ {
			rb.Enqueue(i)
		}
		for i := 0; i < m; i++ {
			_, _ = rb.Dequeue()
		}
	}
}

func BenchmarkEnqueue(b *testing.B) {
	rb := New[int](2 * 1024 * 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rb.Enqueue(i)
	}
}

func BenchmarkDequeue(b *testing.B) {
	rb := New[int](2 * 1024 * 1024)
	rb.n = len(rb.a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rb.Dequeue()
	}
}

func BenchmarkPeek(b *testing.B) {
	rb := New[int](2 * 1024 * 1024)
	rb.n = len(rb.a)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rb.Peek(i)
	}
}
