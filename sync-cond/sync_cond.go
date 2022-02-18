package main

import (
	"fmt"
	"sync"
	"time"
)

/* a rendezvous point for goroutines waiting for or announcing the occurence of
an event
*/

func main() {
	mutex := &sync.Mutex{}
	cond := sync.NewCond(mutex)

	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		cond.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		cond.L.Unlock()
		cond.Signal()
	}

	for i := 0; i < 10; i++ {
		cond.L.Lock()

		for len(queue) == 2 {
			cond.Wait()
		}

		fmt.Println("Adding to queue")

		queue = append(queue, struct{}{})

		go removeFromQueue(1 * time.Second)

		cond.L.Unlock()
	}
}
