package handy

import (
	"fmt"
	"math/rand"
)

func TakeFirstN(doneChan <-chan interface{}, inputStream <-chan interface{}, N int) <-chan interface{} {
	results := make(chan interface{})

	go func() {
		defer close(results)

		for i := 0; i < N; i++ {
			select {
			case <-doneChan:
				return
			default:
				results <- (<-inputStream)
			}
		}
	}()

	return results
}

func RepeatFn(doneChan <-chan interface{}, fn func() interface{}) <-chan interface{} {
	results := make(chan interface{})

	go func() {
		for {
			select {
			case <-doneChan:
				return
			default:
				results <- fn()
			}
		}
	}()

	return results
}

func main() {

	doneChan := make(chan interface{})

	defer close(doneChan)

	rand := func() interface{} {
		return rand.Int()
	}

	inputStream := RepeatFn(doneChan, rand)

	for num := range TakeFirstN(doneChan, inputStream, 10) {
		fmt.Println(num)
	}
}
