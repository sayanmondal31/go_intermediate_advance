package main

import "fmt"

func producers(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)
}

func filter(in <-chan int, out chan<- int) {
	for val := range in {
		if val%2 == 0 {
			out <- val
		}
	}

	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producers(ch1)
	go filter(ch1, ch2)

	for val := range ch2 {
		fmt.Println(val)
	}
}

// //--------------- RANGE OVER CLOSED CHANNEL
// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for val := range ch {
// 		fmt.Println("Received:", val)
// 	}
// }

// //---------- RECEIVING FROM CLOSED CHANNEL
// func main() {
// 	ch := make(chan int)
// 	close(ch)

// 	val, ok := <-ch

// 	if !ok {
// 		fmt.Println("Channel is closed")
// 	} else {
// 		fmt.Println(val)
// 	}

// }

// // ------ SIMPLE CLOSING CHANNEL
// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}

// 		close(ch)
// 	}()

// 	for value := range ch {
// 		fmt.Println(value)
// 	}
// }
