package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var writeOps uint64

	N := 1
	C := 1000000

	start := time.Now()

	wg := &sync.WaitGroup{}

	for i := 0; i < N; i++ {

		wg.Add(1)

		go func() {
			for j := 0; j < C; j++ {
				atomic.AddUint64(&writeOps, 1)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("go-routines:", N, " each-loop:", C, " duration:", time.Since(start), " res:", writeOps)
}
