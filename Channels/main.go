package main

import "fmt"

func main() {

	// variable := make(chan type)
	greeting := make(chan string)
	errch := make(chan error)

	greetString := "Hello"
	go func() {
		greeting <- greetString //receive : towards the channel ; blocking because it is continously trying to receive
		// values, it is ready to receive continous flow of data

	}()

	go errorFu(errch)

	receiver := <-greeting // why it's not blocking because it's in main thread

	fmt.Printf("%v\n", receiver)

	err := <-errch
	fmt.Printf("this is error: %s\n", err)

	fmt.Println("End of program")
}

func errorFu(errch chan error) {
	er := fmt.Errorf("error channnel")

	errch <- er

	<-errch
}
