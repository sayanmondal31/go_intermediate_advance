package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	tokens     chan struct{} //empty structs is actually zero memory overheads
	refillTime time.Duration // how often new tokens added to bucket
}

// params : rateLimit - max num of actions in specified time frame
// params: refillTime - after which time new token will be filled
func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:     make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}

	for range rateLimit { // first time to refill when initialize
		rl.tokens <- struct{}{} // send imaginary data zero memory to token channels
	}

	go rl.startRefill() // start to refill

	return rl // return with filled tokens
}

func (rl *RateLimiter) startRefill() {
	ticker := time.NewTicker(rl.refillTime) //how much time to refill
	defer ticker.Stop()                     // anything exits or stops , stop ticker to refill

	for { // loop to keep up running to refill
		select {
		case <-ticker.C: //when ticker sends signal to complete
			select {
			case rl.tokens <- struct{}{}: //new token will be added after re
			default:
			}
		}
	}

}

func (rl *RateLimiter) allow() bool {
	select {
	case <-rl.tokens: // now tokens are removed from channels to perform operations
		return true
	default:
		return false
	}
}

func main() {
	rateLimiter := NewRateLimiter(5, time.Second) // allows 5 reqs

	for range 10 {
		if rateLimiter.allow() {
			fmt.Println("Request Allowed")
		} else {
			fmt.Println("Request denied")
		}
		time.Sleep(5 * time.Millisecond)
	}
}
