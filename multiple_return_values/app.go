package main

import "fmt"

// Func signature indicates that we are returning 2 integers
func vals() (int, int) {
	return 3, 7
}

func main() {
	// Use both return values with multiple assignment, fancy
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// _ is the blank identifier, I guess it is used for when we don't want to use
	// a certain return value
	_, c := vals()
	fmt.Println(c)
}