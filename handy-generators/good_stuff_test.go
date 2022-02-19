package main

import (
	"math/rand"
	"testing"
)

func BenchmarkStuff(b *testing.B) {
	doneChan := make(chan interface{})
	defer close(doneChan)

	b.ResetTimer()

	rand := func() interface{} {
		return rand.Int()
	}

	inputStream := RepeatFn(doneChan, rand)

	for range TakeFirstN(doneChan, inputStream, b.N) {

	}
}

/*
goos: windows
goarch: amd64
pkg: gcp/handy-generators
cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
BenchmarkStuff
BenchmarkStuff-12
 1465440               822.3 ns/op             8 B/op          1 allocs/op
PASS
ok      gcp/handy-generators    2.191s
*/
