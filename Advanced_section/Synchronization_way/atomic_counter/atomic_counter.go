package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	count int64
}

func (ac *AtomicCounter) increament() {
	atomic.AddInt64(&ac.count, 1) // &ac.count - actual value that needs to incremented
}

func (ac *AtomicCounter) getValue() int64 {
	return atomic.LoadInt64(&ac.count) // reading data safely without any interruption
}

func main() {

	var wg sync.WaitGroup

	numWorkers := 10
	wg.Add(numWorkers)
	// value := 1

	counter := &AtomicCounter{}

	for range numWorkers {
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increament()
				// value++
			}

		}()
	}

	wg.Wait()

	fmt.Println("Final counter value: ", counter.getValue())
	// fmt.Println("Final counter value: ", value)

}
