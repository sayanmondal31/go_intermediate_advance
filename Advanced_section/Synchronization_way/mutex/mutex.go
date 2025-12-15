package main

import (
	"fmt"
	"sync"
)

// // ---------------- MUTEX WITHOUT STRUCT (MOSTLY USED) >>>>>>>>>>>>>>>>>>>>>>
func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numWorkers := 5
	wg.Add(numWorkers)

	increament := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()   // marks the beginning critical section
			counter++   // protected
			mu.Unlock() // marks the end critical section
		}

	}

	for range numWorkers {
		go increament()
	}
	wg.Wait()
	fmt.Println("Final counter value: ", counter)

}

// // ---------------- MUTEX WITH STRUCT (MOSTLY USED) >>>>>>>>>>>>>>>>>>>>>>
// // mutex -  this ensure multiple go routines safely increament without interfaring with each other

// type Counter struct {
// 	mu    sync.Mutex //handle locking, only one
// 	count int
// }

// func (c *Counter) increament() {
// 	c.mu.Lock()         // this locks so that other go routines can't access counter
// 	defer c.mu.Unlock() // releases on success or fails
// 	c.count++
// }

// func (c *Counter) getValue() int { // this locks Counter(struct) before accessing value and after the operation done
// 	// this ensure thread safe go routines
// 	c.mu.Lock()
// 	defer c.mu.Unlock()

// 	return c.count
// }

// func main() {
// 	var wg sync.WaitGroup // initiate wait group

// 	counter := &Counter{} // creates instance and return pointer of Counter
// 	numWorkers := 10      // setting number of workers

// 	// wg.Add(numWorkers)

// 	for range numWorkers {
// 		wg.Add(1) // each time we initiate the go routine
// 		go func() {
// 			defer wg.Done() // once go routine is finished then subtract from wait group
// 			for range 1000 {
// 				counter.increament()
// 				// counter.count++
// 			}
// 		}()
// 	}

// 	wg.Wait()

// 	fmt.Printf("final counter value: %d\n", counter.getValue())

// }

// // ---------------- MUTEX WITH STRUCT (MOSTLY USED) <<<<<<<<<<<<<<<<<<<<<
