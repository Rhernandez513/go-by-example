package main

import "fmt"
import "time"

func main() {

	// simulate some requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	//This limiter channel will receive a value every 200 milliseconds. This is the regulator in our rate limiting scheme.``
	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		//By blocking on a receive from the limiter channel before serving each request
		//we limit ourselves to 1 request every 200 milliseconds.
		<-limiter
		// process the request by printing it to the screen
		fmt.Println("request", req, time.Now())
	}

	// We can allow for short bursts of request handling by creating a buffered channel
	burstyLimiter := make(chan time.Time, 3)

	//Fill up the channel to represent allowed bursting.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	//Every 200 milliseconds we’ll try to add a new value to burstyLimiter, up to its limit of 3.
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<- burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

