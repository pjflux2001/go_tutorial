package main

import (
	"fmt"
	"sync"
)

// since go was born in world of multi-core and multi-thread world
// channels allow data between multiple go routines
// to avoid race conditions and memory issues

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		// direction of data flow : channel to i
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		// direction of data flow : to channel
		ch <- 42
		wg.Done()
	}()
	wg.Wait()
}
