package main

import (
	"fmt"
	"time"
)

// -------------- ticker with time to stop
func main() {
	ticker := time.NewTicker(time.Second)
	stop := time.After(5 * time.Second)

	defer ticker.Stop()

	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("tick at:", tick)
		case <-stop:
			fmt.Println("Stopping ticker")
			return
		}
	}
}

// -------------- SCHEDULING LOGGING, PERIODIC TASKS, POLLING FOR UPDATES
// func periodicTask() {
// 	fmt.Println("Performing perofic task at: ", time.Now())
// }

// // used for scheduling, periodic tasks, polling
// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}

// }

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	// for tick := range ticker.C {
// 	// 	fmt.Println("Tick at:", tick)
// 	// }

// 	i := 1
// 	// for range 5 {
// 	// 	i *= 2
// 	// 	fmt.Println(i)

// 	// }

// 	for range ticker.C {
// 		i *= 2
// 		fmt.Println(i)

// 	}

// }
