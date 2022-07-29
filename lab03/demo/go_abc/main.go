package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// go say("world")
	// say("hello")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		//do stuff
		wg.Done()
	}()
	wg.Wait()
}

/*
Timeline 1 -> all good
th2 goroutine
wg.Add(1)

th1 main
wg.Wait()

th2 
wg.Done()

Timeline 2 -> wrong
th 1 main
wg.Wait()

th 2 goroutine
wg.Add()

both are possible if add is in the goroutine
*/
