package main

import (
	"fmt"
	"sync"
	"time"
)

func RecvLeaky() {
	doWork := func(strChan <-chan string, wg *sync.WaitGroup) <-chan interface{} {
		res := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited")
			defer wg.Done()
			defer close(res)

			for str := range strChan {
				res <- str
			}

		}()

		return res
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)

	doWork(nil, wg)

	wg.Wait()
}

func RecvNotLeaky() {
	doWork := func(strChan <-chan string, wg *sync.WaitGroup, doneChan <-chan interface{}) <-chan interface{} {
		resChan := make(chan interface{})

		go func() {
			defer wg.Done()
			defer fmt.Println("doWork exited")
			defer close(resChan)

			for {
				select {
				case <-doneChan:
					return
				case val := <-strChan:
					resChan <- val
				}
			}
		}()

		return resChan
	}

	doneChan := make(chan interface{})

	wg := &sync.WaitGroup{}

	wg.Add(1)

	resChan := doWork(nil, wg, doneChan)

	go func() {
		time.Sleep(2 * time.Second)

		fmt.Println("Cancelling doWork ")

		close(doneChan)
	}()

	fmt.Println(<-resChan)

	wg.Wait()
}

func main() {
	// RecvLeaky()

	RecvNotLeaky()
}
