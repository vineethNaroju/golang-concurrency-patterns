package main

import "fmt"

func Adhoc() {
	data := make([]int, 4)

	loopData := func(dataChan chan<- int) {
		defer close(dataChan)

		for _, val := range data {
			dataChan <- val
		}
	}

	dataChan := make(chan int)

	go loopData(dataChan)

	for x := range dataChan {
		fmt.Println(x)
	}
}

func Confined() {

	data := make([]int, 4)

	readOnlyChan := func() <-chan int {
		dataChan := make(chan int)

		go func() {

			defer close(dataChan)

			for _, val := range data {
				dataChan <- val
			}
		}()

		return dataChan
	}

	consumer := func(readOnlyChan <-chan int) {
		for val := range readOnlyChan {
			fmt.Println(val)
		}
	}

	dataChan := readOnlyChan()

	consumer(dataChan)
}

func main() {

}
