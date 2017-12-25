package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
	fmt.Println(from, ":", i)
	}
}

func main() {
	// synchronous execution
	f("direct")

	// concurrent execution
	go f("goroutine")

	// anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// this line requires that we hit enter to complete program execution
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

}
