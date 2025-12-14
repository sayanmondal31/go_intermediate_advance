package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func doWork(ctx context.Context) { // this is infinite loop, but return on specific case
	for { // running infinite loop
		select {
		case <-ctx.Done(): // if context deadline occured
			fmt.Println("Work cancelled", ctx.Err())
			return
		default:
			fmt.Println("Working")
		}
		// Setting speed loop every 500ms
		time.Sleep(500 * time.Millisecond) // after half second it call again
	}

}

func main() {
	rootCtx := context.Background() // root context, just a placeholder

	// ctx, cancel := context.WithTimeout(rootCtx, 2*time.Second) // timer of the context starts here, have time duration 2 sec
	// After this time duration context send cancel signal

	ctx, cancel := context.WithCancel(rootCtx)

	go func() {
		time.Sleep(2 * time.Second) // simulating heavy task. Consider this a heavy time consuming
		cancel()                    // manual cancel after task complete
	}()

	// defer cancel() // calls right before the main function exits, automatic cancel

	ctx = context.WithValue(ctx, "requestID", "joh234234")

	go doWork(ctx) // this runs in background
	time.Sleep(3 * time.Second)

	requestID := ctx.Value("requestID") // after the time limit context still exist

	if requestID != nil {
		fmt.Println("Request ID:", requestID)
	} else {
		fmt.Println("No request id found")
	}

	logWithContext(ctx, "This is test log message")
}

func logWithContext(ctx context.Context, message string) {
	reqVal := ctx.Value("requestID")
	log.Printf("Request Id: %v - %v ", reqVal, message)
}

// func checkEventAndODD(ctx context.Context, num int) string {
// 	select {
// 	case <-ctx.Done(): // when the cancellation signal is received
// 		return "Operatio Cancelled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("%d is even", num)
// 		} else {
// 			return fmt.Sprintf("%d is odd ()", num)
// 		}
// 	}
// }

// func main() {
// 	ctx := context.TODO() //If we don't know which context we have to choose

// 	result := checkEventAndODD(ctx, 5)
// 	fmt.Println("result with context.TODO", result)

// 	ctx = context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, 1*time.Second) // this context will send cancel signal after 1 second

// 	defer cancel()

// 	result = checkEventAndODD(ctx, 10)
// 	fmt.Println("result with context.WithTimeout", result)

// 	time.Sleep(3 * time.Second)
// 	result = checkEventAndODD(ctx, 15)
// 	fmt.Println("reult after timeout", result)
// }

// // ================= DIFFERENCE BETWEEN CONTEXT.TODO AND CONTEXT.BACKGROUND
// func main() {
// 	todoContext := context.TODO() // placeholder
// 	contextBkg := context.Background()

// 	ctx := context.WithValue(todoContext, "name", "john")

// 	fmt.Println(ctx)
// 	fmt.Println(ctx.Value("name"))

// 	ctx = context.WithValue(contextBkg, "city", "New york")
// 	fmt.Println(ctx)
// 	fmt.Println(ctx.Value("city"))

// }
