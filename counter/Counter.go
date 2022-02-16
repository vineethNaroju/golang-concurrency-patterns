package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mutex sync.Mutex
	mp    map[string]int
}

func (c *Counter) Inc(key string) {
	c.mutex.Lock()
	c.mp[key]++
	c.mutex.Unlock()
}

func (c *Counter) Get(key string) int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.mp[key]
}

func main() {

	counter := &Counter{
		mp: make(map[string]int),
	}
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(c *Counter, wg *sync.WaitGroup) {
			defer wg.Done()
			c.Inc("a")
		}(counter, wg)
	}

	wg.Wait()

	fmt.Println(counter.Get("a"))
}
