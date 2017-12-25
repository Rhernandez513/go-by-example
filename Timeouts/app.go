package main

import "fmt"
import "time"

// implementing timeouts in Go is easy due to Channels and Select
// Running this program shows the first operation timing out and the second succeeding.
func main() {

	c1 := make(chan string, 1)

	// Simulate an external call
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// Timeout only is selected if the full second elapses
	select {
		case res := <- c1:
			fmt.Println(res)
		case <- time.After(time.Second * 1)	:
			fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
		case res := <- c2:
			fmt.Println(res)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout 2")
	}
}
