package main

import (
	"fmt"
	"time"
)

func main() {

	// variable := make(chan type)
	greeting := make(chan string)

	greetString := "Hello"
	go func() {
		greeting <- greetString //receive : towards the channel ; blocking because it is continously trying to receive
		// values, it is ready to receive continous flow of data
		greeting <- "world"

		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}

	}()
	// go func() {
	// 	receiver := <-greeting
	// 	fmt.Println(receiver)
	// 	receiver = <-greeting
	// 	fmt.Println(receiver)
	// }()
	receiver := <-greeting
	fmt.Println(receiver)
	receiver = <-greeting
	fmt.Println(receiver)

	for range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}

	time.Sleep(time.Second)

	fmt.Println("End of program")

}
