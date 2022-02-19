package main

import "testing"

func BenchmarkPipeline(b *testing.B) {
	doneChan := make(chan interface{})
	defer close(doneChan)

	for i := 0; i < b.N; i++ {
		inputStream := Generator(doneChan, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 23, 14, 15)
		multipliedStream := Multiply(doneChan, inputStream, 10)
		addedStream := Add(doneChan, multipliedStream, 200)
		for range addedStream {

		}
	}

}

/*
goos: windows
goarch: amd64
pkg: gcp/pipelines
cpu: Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz
BenchmarkPipeline
BenchmarkPipeline-12
   71775             16697 ns/op             546 B/op          7 allocs/op
PASS
ok      gcp/pipelines   1.530s
*/
