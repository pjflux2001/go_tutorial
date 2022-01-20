package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// os threads aren't used as in other languages
// layer of abstraction over threads
func main() {
	// go keyword for things to work in thread
	msg := "1"
	wg.Add(1)
	go func() {
		fmt.Println(msg)
		wg.Done()
	}()
	msg = "2"
	// time.Sleep(100 * time.Millisecond) // bad practice
	// hence using wait group
	wg.Wait() // good practice
}
