// Basic concurrency example for Go
// From https://medium.com/@matryer/very-basic-concurrency-for-beginners-in-go-663e63c6ba07#.1behz6ek8

package main

import (
	"fmt"
	"time"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	for n := 2; n <= 12; n++ {
		// Add a new counter to the WaitGroup
		waitGroup.Add(1)
		// Start a new goroutine
		go timesTable(n)
	}
	// Set the main function to wait before exiting
	waitGroup.Wait()
}

func timesTable(x int) {
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", x, i, x*i)
		time.Sleep(100 * time.Millisecond)
	}
	// Remove a counter from the WaitGroup
	waitGroup.Done()
}