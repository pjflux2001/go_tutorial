package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{} // used globally here
var counter = 0
var m = sync.RWMutex{} // to maintain mutual exclusion
// any reader will wait for all writers to get done
// any writer will wait for all readers to get done

func main() {
	runtime.GOMAXPROCS(100) // setting number of threads
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock() // works
		go sayHello()
		m.Lock() // works
		go increment()
		// number of threads
		fmt.Println("Threads : ", runtime.GOMAXPROCS(-1))
	}
	wg.Wait()
}

func sayHello() {
	// m.RLock() // didn't work as required
	fmt.Println("Hello Cnt : ", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	// m.Lock() // didn't work as required
	counter++
	m.Unlock()
	wg.Done()
}
