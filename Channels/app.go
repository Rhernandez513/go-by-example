package main

// channels are pipes that connect concurrent goroutines
// you can send values into channels from one goroutine and receive
// those values ito another goroutine

import "fmt"

func main() {
	// create a new channel with    make(chan value-type)
	// channels are typed by the values they convey
	messages := make(chan string)

	// Sending a value into a channel using the <- syntax
	// This is done here anonymously
	go func() { messages <- "ping"} ()

	// receive a value from a channel
	msg := <-messages
	fmt.Println(msg)
	// This successfully passes a message from one goroutine to another via our channel

	// by default sends and receives block until bother the sender and receiver are ready
	// this allows us to wait at the end of our program for the "ping" message without having to
	// use any other form of synchronization
}