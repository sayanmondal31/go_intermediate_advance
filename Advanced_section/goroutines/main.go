package main

import (
	"fmt"
	"time"
)

// Go routines are just functions tht leave the main thread and run in the background and come back to join the
// main thread once the function exceution is done to return any value

// Go routines don't stop the program flow and non blocking, like async and await in js

func main() {
	var err error
	fmt.Println("Begining program")
	go sayHello() // non blocking mechanism
	fmt.Println("After sayHello function")
	go func() {
		err = doWork()
	}()
	go printNumbers()
	go printLetters()

	time.Sleep(2 * time.Second) // waits for 2 seconds -> blocks
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Work complete successfully")
	}

}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from go routine")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Numner: ", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	time.Sleep(1 * time.Second)

	return fmt.Errorf("an error occured in dowork")
}
