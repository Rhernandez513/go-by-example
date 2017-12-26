package main

import "fmt"
import "time"

// Timers are for when you want to do something once in the future
// tickers are for when you want to do something repeatedly at regular intervals.
func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	// Tickers use a similar mechanism to timers: a channel that is sent values.
	// Here we’ll use the range builtin on the channel to iterate over the values as they arrive every 500ms.
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
