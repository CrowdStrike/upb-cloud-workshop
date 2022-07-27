package main

import (
	"fmt"
	// "sync"
	// "time"
)
// var wg sync.WaitGroup
// var mu sync.Mutex
var v  map[string]int


func inc(key string) {
	v[key]++
}

func getValue(key string) int {
	return v[key]
}

func main() {	
	v = map[string]int{}

	for i := 0; i < 24; i++ {
		go inc("somekey")
	}

	// time.Sleep(time.Second)
	fmt.Println(getValue("somekey"))	
}
