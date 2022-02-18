package main

import (
	"fmt"
	"sync"
)

type Button struct {
	Clicked *sync.Cond
}

func main() {
	button := &Button{
		Clicked: sync.NewCond(&sync.Mutex{}),
	}

	subscribe := func(cond *sync.Cond, cb func()) {
		vg := &sync.WaitGroup{}
		vg.Add(1)

		// to make sure below goroutine actually runs and not scheduled later

		go func() {
			vg.Done()
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait()
			cb()
		}()

		vg.Wait()
	}

	// just to wait until our code prints stuff
	wg := &sync.WaitGroup{}

	wg.Add(1)
	subscribe(button.Clicked, func() {
		fmt.Println("Yo-1")
		wg.Done()
	})

	wg.Add(1)
	subscribe(button.Clicked, func() {
		fmt.Println("Yo-2")
		wg.Done()
	})

	button.Clicked.Broadcast()

	wg.Wait()
}
