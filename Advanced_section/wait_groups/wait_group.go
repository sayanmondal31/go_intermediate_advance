package main

import (
	"fmt"
	"sync"
	"time"
)

// WAIT GROUP VERSION OF TICKETING SYSTEM >>>>>>>>>>>>>>>>>>
type Ticket struct {
	PersonID   int
	NumTickets int
	Cost       int
}

func ProcessTicketsWorker(wg *sync.WaitGroup, requests <-chan Ticket, result chan<- int) {
	defer wg.Done()
	for request := range requests {
		time.Sleep(time.Second) //simulating ticket process
		// fmt.Printf("Ticket processed for PersonID %d", request.PersonID)
		result <- request.PersonID
	}

}

func main() {
	var wg sync.WaitGroup
	numRequests := 5
	numWorkers := 3
	requests := make(chan Ticket, numRequests)
	results := make(chan int)

	wg.Add(numWorkers)

	for range numWorkers {
		go ProcessTicketsWorker(&wg, requests, results)
	}

	for i := range numRequests {
		priceOfEachTicket := 5
		numTickets := 2
		requests <- Ticket{PersonID: i + 1, NumTickets: numTickets, Cost: numTickets * priceOfEachTicket}
	}
	close(requests)

	go func() {

		wg.Wait()
		close(results)

	}()

	for result := range results {
		fmt.Printf("✅Ticket processed for PersonId: %d\n", result)
	}

}

// WAIT GROUP VERSION OF TICKETING SYSTEM >>>>>>>>>>>>>>>>>>

// // ----------------CONSTRUCTION EXAMPLE >>>>>>>>>>>>>>>>>>
// type Worker struct {
// 	ID    int
// 	Tasks string
// }

// // performTask simuates a worker performing a task
// func (w *Worker) PerformTask(wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	fmt.Printf("WorkerID %d started %s\n", w.ID, w.Tasks)
// 	time.Sleep(time.Second)
// 	fmt.Printf("WorkerID %d finished %s\n", w.ID, w.Tasks)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	// tasks to be performed by worker
// 	tasks := []string{"digging", "laying brics", "painting"}

// 	for i, task := range tasks {
// 		worker := Worker{ID: i + 1, Tasks: task}
// 		wg.Add(1) // this add 1 by 1
// 		go worker.PerformTask(&wg)
// 	}

// 	// wait for all workers to complete
// 	wg.Wait()

// 	// construction is finished
// 	fmt.Println("Construction is finished")
// }

// // ----------------CONSTRUCTION EXAMPLE >>>>>>>>>>>>>>>>>>

// //------------- WAIT GROUPS with channels>>>>>>>>>>>>>>>>>>
// func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("WorkerID %d starting!\n", id)
// 	time.Sleep(time.Second) //simulate worlk
// 	for task := range tasks {

// 		results <- task * 2
// 	}
// 	fmt.Printf("WorkerID %d Stopped!\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	numWorkers := 3
// 	numjobs := 5
// 	results := make(chan int, numjobs)
// 	tasks := make(chan int, numjobs)

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i+1, tasks, results, &wg)
// 	}

// 	for i := range numjobs { // data is sent to channel
// 		tasks <- i
// 	}
// 	close(tasks)

// 	go func() {
// 		wg.Wait()
// 		close(results)

// 	}()

// 	for result := range results { // receiving channel data
// 		fmt.Println("result:", result)
// 	}

// }

// //------------- WAIT GROUPS with channels>>>>>>>>>>>>>>>>>>

// // BASIC WAIT GROUPS without channels>>>>>>>>>>>>>>>>>>>>>>
// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done() //decrement count by 1
// 	// wg.Add(1) // WRONG PRACTICE
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second) // simulate some time spent on processing the task

// 	fmt.Printf("Worker %d finished\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3

// 	wg.Add(numWorkers) // wait for 3 call of waitgroup.done; GOOD PRACTICE

// 	// launch workers
// 	startTime := time.Now()
// 	for i := range numWorkers {

// 		go worker(i, &wg)

// 	}

// 	wg.Wait() // this will wait all the go routines to complete; Blocking
// 	endtime := time.Now()
// 	fmt.Printf("All workers finished in %.2f s \n", endtime.Sub(startTime).Seconds())

// }

// mandlebrot test
// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// 	"time"
// )

// // BASIC WAIT GROUPS without channels>>>>>>>>>>>>>>>>>>>>>>

// // The maximum number of iterations to check for a point's divergence. >>>>>>>>>>>>>>>>
// const MaxIterations = 10000

// // Function to calculate the number of iterations required for a single complex number (c)
// func mandelbrot(c complex128) int {
// 	z := complex128(0)
// 	i := 0
// 	// Loop until the magnitude of z exceeds 2 or we hit MaxIterations
// 	for ; i < MaxIterations && real(z)*real(z)+imag(z)*imag(z) < 4; i++ {
// 		z = z*z + c
// 	}
// 	return i
// }

// // Worker function that calculates a segment of the Mandelbrot Set
// func worker(start int, end int, size int, wg *sync.WaitGroup, results []int) {
// 	defer wg.Done()

// 	// Define the complex plane region to calculate (e.g., a specific row or chunk)
// 	const xMin, xMax = -2.0, 1.0
// 	const yMin, yMax = -1.5, 1.5

// 	// The core work loop
// 	for i := start; i < end; i++ {
// 		// Map the linear index 'i' back to (x, y) coordinates on the grid
// 		x := float64(i % size)
// 		y := float64(i / size)

// 		// Scale the coordinates to the complex plane
// 		re := xMin + x*(xMax-xMin)/float64(size)
// 		im := yMin + y*(yMax-yMin)/float64(size)

// 		c := complex(re, im)

// 		// Perform the complex, iterative calculation
// 		results[i] = mandelbrot(c)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	// 1. Define the total work size
// 	const imageSize = 4000               // 4000x4000 resolution
// 	totalPoints := imageSize * imageSize // Total work is 16 million points

// 	// 2. Determine the number of concurrent tasks (using your 12 cores + 4 for buffer)
// 	numWorkers := runtime.NumCPU() * 4
// 	// numWorkers will likely be 12 * 4 = 48 on your M3 Pro

// 	chunkSize := totalPoints / numWorkers

// 	// Pre-allocate the result array to store calculation results
// 	results := make([]int, totalPoints)

// 	fmt.Printf("Starting CPU-bound task:\n")
// 	fmt.Printf("Total Points to Calculate: %d\n", totalPoints)
// 	fmt.Printf("Number of Cores Available: %d\n", runtime.NumCPU())
// 	fmt.Printf("Number of Goroutines Used: %d\n", numWorkers)

// 	// Launch workers
// 	startTime := time.Now()
// 	wg.Add(numWorkers)

// 	for w := 0; w < numWorkers; w++ {
// 		start := w * chunkSize
// 		end := (w + 1) * chunkSize
// 		if w == numWorkers-1 {
// 			end = totalPoints // Ensure the last worker handles any remainder
// 		}

// 		// Each worker gets a distinct chunk of the calculation
// 		go worker(start, end, imageSize, &wg, results)
// 	}

// 	// Wait for all calculation to finish
// 	wg.Wait()
// 	endTime := time.Now()

// 	duration := endTime.Sub(startTime)

// 	fmt.Printf("\n--- Results ---\n")
// 	fmt.Printf("Calculation finished in %v\n", duration)
// }
// // The maximum number of iterations to check for a point's divergence. >>>>>>>>>>>>>>>>
