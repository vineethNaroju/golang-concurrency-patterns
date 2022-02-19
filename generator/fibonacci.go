package main

import "fmt"

func Fib(N int) <-chan int {
	c := make(chan int)

	go func() {

		for a, b := 0, 1; a < N; a, b = a+b, a {
			c <- a
		}

		close(c)
	}()

	return c
}

func main() {

	for x := range Fib(20) {
		fmt.Println(x)
	}

}
