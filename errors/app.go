package main

import "errors"
import "fmt"

// Basic error
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg int
	prob string
}

// possible to use custom types are errors by implementing the Error() func on them
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %2", e.arg, e.prob)
}

func f2 (arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		// Looks like
		// if return value, error, and error is not null

		// Tutorial says: "the use of an inline error check on the if line is a common idiom in Go code"
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	// discard return value, e is the error
	_, e := f2(42)

	// return value is the struct, ok is the error in the form of a bool?

	// Tutorial says:
	// If you want to programmatically use the data in a custom error, you'll need to get the error as an instance
	// of the custom type error via type assertion
	//
	// So it looks like this is an example of type assertion in Go?
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}