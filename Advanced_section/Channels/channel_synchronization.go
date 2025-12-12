package main

// func main() {
// 	done := make(chan struct{}) // unbuffered chan

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(2 * time.Second)
// 		done <- struct{}{}
// 	}()

// 	<-done // emit from done chan
// 	fmt.Println("Finished")
// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		ch <- 9 // Blocking until the value is received
// 		time.Sleep(time.Second)
// 		fmt.Println("Sent value")
// 	}()

// 	value := <-ch // Blocking until the value is sent

// 	fmt.Println(value)
// }

// // ---------------------------- CHANNEL SYNCHRONIZATION OF MULTIPLE GO ROUTINES------------------
// func main() {
// 	numGoroutines := 3
// 	done := make(chan int, 3)

// 	for i := range numGoroutines { // number of goroutines is should be equal to same number of receivers
// 		// time.Sleep(2 * time.Second)
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working...\n", id)
// 			done <- id // sending signa of completion
// 		}(i)
// 	}

// 	for range numGoroutines { // number of goroutines is should be equal to same number of receivers
// 		<-done // Wait for each goroutines to finish, wait for all goroutined to signal completion
// 	}

// 	fmt.Println("All goroutines are finished")
// }

// ------------------------- SYNCHRONIZing  DATA EXCHANGE
// func main() {
// 	// continous flow chatting data, weather data, any stream data...
// 	data := make(chan string)

// 	go func() {
// 		for i := range 5 {
// 			data <- "hello " + string('0'+i)
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 		close(data) // close channel
// 	}()

// 	// close(data) // Channel closed before Goroutine could send a value to channel

// 	for value := range data { // recevier
// 		fmt.Println("Received value: ", value, ":", time.Now())
// 	} // Loops over only on active channel, creates receiver each time and stops receiver (looping)  once the channel id closed
// }
