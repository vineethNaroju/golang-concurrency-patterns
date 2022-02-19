package main

import "testing"

func BenchmarkRib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for range Fib(i) {

		}
	}
}

/*
goos: windows
goarch: amd64
pkg: gcp/generator
cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
BenchmarkRib-12    	  200011	      7728 ns/op	     120 B/op	       2 allocs/op
PASS
ok  	gcp/generator	1.772s
*/
