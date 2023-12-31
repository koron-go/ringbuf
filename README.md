# koron-go/ringbuf

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron-go/ringbuf)](https://pkg.go.dev/github.com/koron-go/ringbuf)
[![Actions/Go](https://github.com/koron-go/ringbuf/workflows/Go/badge.svg)](https://github.com/koron-go/ringbuf/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron-go/ringbuf)](https://goreportcard.com/report/github.com/koron-go/ringbuf)

See <https://kumagi.hatenablog.com/entry/ring-buffer>

## Variations

Type         | Desc.
-------------|-------
RingBuffer0  | Naive implementation with write and read indexes with modulo
RingBuffer1  | Use bit **and** op instead of modulo (based RingBuffer0)
RingBuffer2  | Support goroutine with `sync.Mutex` (based RingBuffer1)
RingBuffer3  | Use sync/atomic for multi-goroutine (based RingBuffer1)
RingBuffer3B | Same as RingBuffer3 but w/o index caches
RingBuffer4  | Use chan for multi-goroutine
RingBuffer4B | Use chan for multi-goroutine with blocking
RingBuffer5  | Use read index and capacity length w/o multi-goroutine supports

## Benchmark results

```console
$ go test -bench . -benchmem
goos: windows
goarch: amd64
pkg: github.com/koron-go/ringbuf
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkRingBuffer0/single-16          177274972                6.697 ns/op           0 B/op          0 allocs/op
BenchmarkRingBuffer1/single-16          835998076                1.431 ns/op           0 B/op          0 allocs/op
BenchmarkRingBuffer2/single-16          84422025                13.75 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer2/multi-16           57571975                22.23 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer3/single-16          135313495                8.905 ns/op           0 B/op          0 allocs/op
BenchmarkRingBuffer3/multi-16           44504277                22.79 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer3B/single-16         125987650                9.456 ns/op           0 B/op          0 allocs/op
BenchmarkRingBuffer3B/multi-16          52267851                23.10 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer4/single-16          62735503                19.07 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer4/multi-16           60012002                20.86 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer4B/single-16         65714535                18.41 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer4B/multi-16          61101039                20.01 ns/op            0 B/op          0 allocs/op
BenchmarkRingBuffer5/single-16          644982705                1.824 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/koron-go/ringbuf     21.503s
```

* 複数goroutine環境での評価
    * sync.Mutex(2)は意外と安定
    * sync/atomic(3)は性能劣化が著しい
    * chan(4)で良いのでは?
* 1 goroutineでの評価
    * index+書き込まれ量の管理方式(5)は、意外と良い。ただし複数goroutine化の見通しが立たないか、もしくは性能が悪そう。
