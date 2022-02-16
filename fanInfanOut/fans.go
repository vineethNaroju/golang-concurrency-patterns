package main

import "fmt"

func generatePipeline(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, val := range nums {
			out <- val
		}
		close(out)
	}()

	return out
}

func squareNums(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for val := range in {
			out <- val * val
		}
		close(out)
	}()

	return out
}

func fanIn(cpos, cneg <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for {
			select {
			case a := <-cpos:
				out <- a
			case b := <-cneg:
				out <- b
			}
		}
	}()

	return out
}

func main() {
	nums := []int{2, 3, 4, 5, 6, 7}

	inputChannel := generatePipeline(nums)

	cpos := squareNums(inputChannel)
	cneg := squareNums(inputChannel)

	c := fanIn(cpos, cneg)

	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += <-c
	}

	fmt.Println(sum)

}
