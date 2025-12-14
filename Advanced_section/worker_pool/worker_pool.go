package main

import (
	"fmt"
	"time"
)

//--------------- BASIC SETUP FOR WORKER POOLS -----------------
// // worker function
// // <- chan int : receive only channel
// // chan <- int : sent only channel
// func worker(id int, tasks <-chan int, results chan<- int) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d processing taks %d\n", id, task)
// 		// simulate some work
// 		time.Sleep(time.Second)
// 		results <- task * 2
// 	}
// }

// func main() {
// 	numWorkers := 3
// 	numJobs := 10
// 	tasks := make(chan int, numJobs)   // buffered channel
// 	results := make(chan int, numJobs) // buffered channel

// 	// step1: create workers
// 	for i := range numWorkers { // creating 3 workers
// 		go worker(i, tasks, results)
// 	}

// 	// step2: send values to the tasks channel
// 	for i := range numJobs { // pass 10 values to tasks channel
// 		tasks <- i
// 	}
// 	close(tasks)

// 	// step3: collect the results
// 	// receive channel should be **|after|** go routine jobs or sent channel
// 	for range numJobs { // receive 10 processed values from go routines
// 		result := <-results
// 		fmt.Println("Result:", result)
// 	}

// }

//--------------- TICKET SELLING AND BUYING IN WORKER POOLS -----------------

type TicketRequest struct {
	PersonID   int
	NumTickets int
	Cost       int
}

// simulate processing of ticket requests
func ticketProcessor(requests <-chan TicketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket(s) of person %d with total cost %d\n", req.NumTickets, req.PersonID, req.Cost)

		// simulate processing time
		time.Sleep(time.Second)

		results <- req.PersonID
	}
}

func main() {
	numRequests := 5
	price := 5
	ticketRequests := make(chan TicketRequest, numRequests)
	ticketResults := make(chan int)

	// start ticket worker(processon)
	// setting 3 workers to work on job queue
	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	// send ticket requests
	for i := range numRequests {

		numTickets := 0

		if i == 0 {
			numTickets = (i + 1) * 2
		} else {
			numTickets = i * 2
		}

		ticketRequests <- TicketRequest{PersonID: i, NumTickets: numTickets, Cost: numTickets * price}
	}
	close(ticketRequests)

	// collect results
	for range numRequests {
		personId := <-ticketResults
		fmt.Printf("🎟️ Ticket for PersonId %d is processed!✅\n", personId)
	}

}
