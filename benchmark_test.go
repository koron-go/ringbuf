package ringbuf

import (
	"testing"
)

func BenchmarkRingBuffer(b *testing.B) {
	rb := New[int](2 * 1024 * 1024)
	b.ResetTimer()
	for n := max(1, b.N/2); n > 0; n -= 1000 {
		m := min(n, 1000)
		for i := 0; i < m; i++ {
			rb.Put(i)
		}
		for i := 0; i < m; i++ {
			_, _ = rb.Get()
		}
	}
}

func BenchmarkPut(b *testing.B) {
	rb := New[int](2 * 1024 * 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rb.Put(i)
	}
}

func BenchmarkGet(b *testing.B) {
	rb := New[int](2 * 1024 * 1024)
	rb.n = b.N
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = rb.Get()
	}
}
