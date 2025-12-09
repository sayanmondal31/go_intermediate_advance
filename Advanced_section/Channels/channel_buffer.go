package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) // initialize channel
	go func() {          // set go routine - separates from main thread
		// ch <- 1 // receiving value
		time.Sleep(2 * time.Second)
		fmt.Println("2 sec goroutine finished")
	}()

	receiver := <-ch // there is receiver and store data to variable

	fmt.Println(receiver, "receiver") // prints
}
