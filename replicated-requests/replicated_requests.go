package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func MakeRequest(doneChan <-chan interface{}, routineId int, wg *sync.WaitGroup, results chan<- int) {
	start := time.Now()

	defer wg.Done()

	timeAfter := time.Duration(2+rand.Intn(10)) * time.Second

	select {
	case <-doneChan:
	case <-time.After(timeAfter):
	}

	select {
	case <-doneChan:
	case results <- routineId:
	}

	taken := time.Since(start)

	if taken < timeAfter {
		taken = timeAfter
	}

	fmt.Println(routineId, " took ", taken)
}

func main() {

	doneChan := make(chan interface{})

	results := make(chan int)

	N := 10

	wg := &sync.WaitGroup{}
	wg.Add(N)

	for i := 0; i < N; i++ {
		go MakeRequest(doneChan, i, wg, results)
	}

	firstOne := <-results
	close(doneChan)
	wg.Wait()

	fmt.Println("Received from : ", firstOne)

}
