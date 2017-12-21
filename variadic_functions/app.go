package main

import "fmt"

// Looks to me that this funciton takes an unlimited number of ints as potential input
func sum(nums ...int) {
	// Looks like the fmt.Print functions' second argument is the separator
	fmt.Print(nums, " ")

}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice,
	// apply them to a variadic function using func(slice...) like this.

	nums := []int{1, 2, 3, 4}
	sum(nums...)
	// I get the creeping feeling that this is an array and not a slice, but it still works                            1
}
