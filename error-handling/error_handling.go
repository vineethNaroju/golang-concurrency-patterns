package main

import (
	"fmt"
	"net/http"
)

func ErrorNotHandled() {
	checkStatus := func(
		doneChan <-chan interface{},
		urls ...string,
	) <-chan *http.Response {
		responsesChan := make(chan *http.Response)

		go func() {
			defer close(responsesChan)

			for _, val := range urls {
				resp, err := http.Get(val)

				if err != nil {
					fmt.Println(err)
					continue
				}

				select {
				case <-doneChan:
					return
				default:
					responsesChan <- resp
				}
			}
		}()

		return responsesChan
	}

	doneChan := make(chan interface{})
	defer close(doneChan)

	urls := []string{"https://www.google.com", "lmao"}

	for resp := range checkStatus(doneChan, urls...) {
		fmt.Println(resp)
	}
}

type Result struct {
	Error    error
	Response *http.Response
}

func ErrorHandled() {
	checkStatus := func(
		doneChan <-chan interface{},
		urls ...string,
	) <-chan Result {
		results := make(chan Result)

		go func() {
			defer close(results)

			for _, val := range urls {
				resp, err := http.Get(val)

				r := Result{Error: err, Response: resp}

				select {
				case <-doneChan:
					return
				default:
					results <- r
				}
			}
		}()

		return results
	}

	doneChan := make(chan interface{})
	defer close(doneChan)

	urls := []string{"https://www.google.com", "oooooh", "https://www.netflix.com"}

	for result := range checkStatus(doneChan, urls...) {
		if result.Error != nil {
			fmt.Printf("error : %v\n\n", result.Error)
			continue
		}

		fmt.Printf("Response: %v\n\n", result.Response)
	}

}

func main() {
	// ErrorNotHandled()

	ErrorHandled()
}
