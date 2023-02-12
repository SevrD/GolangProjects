package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex
	m := make(map[int]int)
	f := func(k int) {
		mutex.Lock()
		m[k] = k
		mutex.Unlock()
	}
	for i := 0; i < 1000; i++ {
		go f(i)

	}
	fmt.Println(m)
}
