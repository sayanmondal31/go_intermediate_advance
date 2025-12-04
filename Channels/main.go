package main

import "fmt"

func main() {

	// variable := make(chan type)
	greeting := make(chan string)

	greetString := "Hello"
	go func() {
		greeting <- greetString //receive : towards the channel ; blocking because it is continously trying to receive
		// values, it is ready to receive continous flow of data

	}()

	receiver := <-greeting // why it's not blocking because it's in main thread

	fmt.Printf("%v\n", receiver)

	fmt.Println("End of program")
}
