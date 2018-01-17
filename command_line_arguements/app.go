package main

import "os"
import "fmt"

func main() {
	argsWithProg := os.Args
	// first value in the slices is the path to the program
	argsWithoutProg := os.Args[1:]

	// normal indexing for individual arguements
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
