package main

import (
	"fmt"
	"time"
)

func TimerTicks() {
	twoSecondChan := make(chan interface{})

	go func() {
		for now := range time.Tick(time.Second * 2) {
			twoSecondChan <- now
		}
	}()

	tenSecondChan := time.After(10 * time.Second)

	for {
		select {
		case val := <-twoSecondChan:
			fmt.Println("every 2 seconds", val)
		case val := <-tenSecondChan:
			fmt.Println("done", val)
			return
		}
	}
}

func main() {

}
