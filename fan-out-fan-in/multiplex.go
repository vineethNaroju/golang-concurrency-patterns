package main

import (
	"fmt"
	"gcp/handy-generators"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func PerfectSquare(doneChan <-chan interface{}, inputStream <-chan interface{}) <-chan interface{} {
	results := make(chan interface{})

	go func() {
		defer close(results)

		for val := range inputStream {

			ok := false

			for i := 0; i*i <= val.(int); i++ {
				if i*i == val {
					ok = true
				}
			}

			select {
			case <-doneChan:
				return
			default:
				if ok {
					results <- val
				}
			}
		}
	}()

	return results
}

func VerySlow() {

	// 21 seconds for finding 10 perfect squares in a stream of random numbers under 5 * 10^9

	rand := func() interface{} {
		return rand.Intn(5000000000) // 5 * 10^9
	}

	doneChan := make(chan interface{})
	defer close(doneChan)

	start := time.Now()

	randIntStream := handy.RepeatFn(doneChan, rand)

	for sq := range handy.TakeFirstN(doneChan, PerfectSquare(doneChan, randIntStream), 10) {
		fmt.Println(sq)
	}

	fmt.Println(time.Since(start))
}

func FanIn(doneChan <-chan interface{}, manyChannels ...<-chan interface{}) <-chan interface{} {
	results := make(chan interface{})

	wg := &sync.WaitGroup{}

	wg.Add(len(manyChannels))

	multiplex := func(inputChan <-chan interface{}) {
		defer wg.Done()

		for val := range inputChan {
			select {
			case <-doneChan:
				return
			default:
				results <- val
			}
		}
	}

	for _, eachChan := range manyChannels {
		go multiplex(eachChan)
	}

	go func() {
		defer close(results)
		wg.Wait()
	}()

	return results
}

func MultiplexedPerfectSquares() {
	rand := func() interface{} {
		return rand.Intn(5000000000)
	}

	doneChan := make(chan interface{})
	defer close(doneChan)

	start := time.Now()

	randIntStream := handy.RepeatFn(doneChan, rand)

	fanOut := runtime.NumCPU()

	parallelFinders := make([]<-chan interface{}, fanOut)

	for i := 0; i < fanOut; i++ {
		parallelFinders[i] = PerfectSquare(doneChan, randIntStream)
	}

	for val := range handy.TakeFirstN(doneChan, FanIn(doneChan, parallelFinders...), 10) {
		fmt.Println(val)
	}

	fmt.Println(time.Since(start))
}

func main() {
	VerySlow()

	fmt.Println("------------------")

	MultiplexedPerfectSquares()
}
