package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	//  if we need one of them >>>>>>>>>
	select {
	case <-timer1.C:
		fmt.Println("timer1 stopped")
	case <-timer2.C:
		fmt.Println("timer 2 stopped")
	}
	//  if we need one of them <<<<<<<<<<

	// if we need both>>>>>>>>>>>>
	// go func() {
	// 	for {
	// 		select {
	// 		case <-timer1.C:
	// 			fmt.Println("timer1 stopped")
	// 		case <-timer2.C:
	// 			fmt.Println("timer 2 stopped")
	// 		}
	// 	}
	// }()

	// time.Sleep(3 * time.Second)
	// if we need both <<<<<<<<<<<<<

}

// // --------------------- Scheduling delayed operation >>>>>>>>>>>>>>>>>>.
// func main() {
// 	timer := time.NewTimer(2 * time.Second)

// 	go func() {
// 		<-timer.C
// 		fmt.Println("Delayed operation executed")
// 	}()

// 	fmt.Println("Waiting")

// 	time.Sleep(3 * time.Second) // blocking timer start
// 	fmt.Println("End of program")
// }

// // --------------------- Timeout operation>>>>>>>>>>>>>>>>>>>>
// func longRunningOps() {
// 	for i := range 20 {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	timeout := time.After(3 * time.Second)
// 	done := make(chan bool)
// 	go func() {
// 		longRunningOps()
// 		done <- true

// 	}()

// 	select {
// 	case <-timeout:
// 		fmt.Println("Operation times out")
// 	case <-done:
// 		fmt.Println("ops Complete")
// 	}
// }
// // --------------------- Timeout operation<<<<<<<<<<<<<<<<<

// --------------------- Basic timer operation
// func main() {
// 	// time.Sleep(time.Second) //time.Sleep -> blocking in nature
// 	fmt.Println("Starting app")
// 	timer := time.NewTimer(2 * time.Second) //time.NewTimer -> Non blocking
// 	fmt.Println("waiting for timer")
// 	stopped := timer.Stop()

// 	if stopped {
// 		fmt.Println("Timer stopped")
// 	}
// 	timer.Reset(time.Second)
// 	fmt.Println("Timer reset")
// 	<-timer.C // blocking in nature,  C is the channel that receives the time
// 	// when timer expired

// 	fmt.Println("timer expired")
// }
