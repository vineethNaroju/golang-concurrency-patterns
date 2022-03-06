package main

import (
	"fmt"
	"sync"
)

func main() {

	cnt := 0

	inc := func() {
		cnt++
	}

	once := &sync.Once{}

	wg := &sync.WaitGroup{}

	N := 10

	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			once.Do(inc)

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(cnt)
}
