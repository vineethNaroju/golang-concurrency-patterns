package main

import "fmt"

func Generator(doneChan <-chan interface{}, vals ...int) <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)

		for _, val := range vals {
			select {
			case <-doneChan:
				return
			default:
				intStream <- val
			}
		}
	}()

	return intStream
}

func Multiply(doneChan <-chan interface{}, inputStream <-chan int, multiplier int) <-chan int {
	resultStream := make(chan int)

	go func() {
		defer close(resultStream)

		for val := range inputStream {
			select {
			case <-doneChan:
				return
			default:
				resultStream <- multiplier * val
			}
		}
	}()

	return resultStream
}

func Add(doneChan <-chan interface{}, inputStream <-chan int, adder int) <-chan int {
	resultStream := make(chan int)

	go func() {
		defer close(resultStream)

		for val := range inputStream {
			select {
			case <-doneChan:
				return
			default:
				resultStream <- (val + adder)
			}
		}
	}()

	return resultStream
}

func main() {
	doneChan := make(chan interface{})

	defer close(doneChan)

	inputStream := Generator(doneChan, 2, 5, 10, 20, -10, 100, 200, 2)

	addedStream := Add(doneChan, inputStream, 100)

	multipliedStream := Multiply(doneChan, addedStream, 3)

	for val := range multipliedStream {
		fmt.Println(val)
	}
}
