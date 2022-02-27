package main

import (
	"fmt"
	"time"
)

/*

Heartbeats lets us know if everything is going fine.

*/

func ExposeHeartbeatAndWork(doneChan <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
	heartbeats := make(chan interface{})
	results := make(chan time.Time)

	go func() {

		defer close(heartbeats)
		defer close(results)

		workPulse := time.Tick(2 * pulseInterval)

		sendPulse := func() {
			select {
			case heartbeats <- struct{}{}:
			default: // No one might be listening to heartbeats, scary
			}
		}

		pulse := time.Tick(pulseInterval)

		sendWork := func(work time.Time) {
			for {
				select {
				case <-doneChan:
					return
				case <-pulse:
					sendPulse()
				case results <- work:
					return
				}
			}
		}

		for {
			select {
			case <-doneChan:
				return
			case <-pulse:
				sendPulse()
			case work := <-workPulse:
				sendWork(work)
			}
		}

	}()

	return heartbeats, results

}

func main() {
	doneChan := make(chan interface{})

	time.AfterFunc(10*time.Second, func() {
		close(doneChan)
	})

	const pulseInterval = 2 * time.Second

	heartbeats, results := ExposeHeartbeatAndWork(doneChan, pulseInterval/2)

	for {
		select {
		case _, ok := <-heartbeats:
			if !ok {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if !ok {
				return
			}
			fmt.Printf("results %v\n", r.Second())
		case <-time.After(pulseInterval):
			return
		}
	}
}
