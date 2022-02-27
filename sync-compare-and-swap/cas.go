package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var x int64 = 10000

	res := atomic.CompareAndSwapInt64(&x, 10000, 897878789)

	fmt.Println(res, x)
}
