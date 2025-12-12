package main

import "fmt"

// ------------------ send only channel
func channel_direction() {
	ch := make(chan int)

	producer(ch)

	consumer(ch)
}

// Send only channel
func producer(ch chan<- int) { // send only
	go func() { // send only channel
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

// Receive only channel
func consumer(ch <-chan int) { // consumer
	for value := range ch {
		fmt.Println("Received: ", value)
	}
}
