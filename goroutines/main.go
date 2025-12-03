package main

import (
	"fmt"
	"time"
)

//

func main() {
	fmt.Println("Begining program")
	go sayHello()
	fmt.Println("After sayHello function")
	time.Sleep(2 * time.Second) // waits for 2 seconds
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from go routine")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is", i)
		time.Sleep(1000 * time.Microsecond)
	}
}
