package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// non-blocking receive
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"

	// non-blocking send
	select {
	case messages <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("no message sent")
	}

	// multi-way non-blocking select with default
	select {
	case msg := <-messages :
		fmt.Println("received message", msg)
	case sig := <- signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}