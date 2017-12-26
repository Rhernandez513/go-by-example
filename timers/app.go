package main

import "time"
import "fmt"

func main() {
	fmt.Println("Waiting for two seconds")
	timer1 := time.NewTimer(time.Second * 2)

	//blocks on the channels' C until it sends a value that indicates the timer
	// has expired
	<-timer1.C
	fmt.Println("two seconds have elapsed")

	timer2 := time.NewTimer(time.Second)

	// an example of stopping a timer before the time has elapsed
	// like a stopwatch!
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer2 stopped")
	}
}
