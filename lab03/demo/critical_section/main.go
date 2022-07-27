package main

import (
	"fmt"
	"sync"
	// "sync"
	// "time"
)

var wg sync.WaitGroup
var mu sync.Mutex
var v map[string]int

func inc(key string) {
	mu.Lock()
	v[key]++
	mu.Unlock()

	wg.Done()
}

func getValue(key string) int {
	return v[key]
}

func main() {
	v = map[string]int{}

	for i := 0; i < 24; i++ {
		wg.Add(1)
		go inc("somekey")
	}

	wg.Wait()
	// time.Sleep(time.Second)
	fmt.Println(getValue("somekey"))
}
