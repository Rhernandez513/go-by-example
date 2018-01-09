package main

import (
	"fmt"
	"os"
)
//import "os"

type point struct {
	x, y int
}

func main()  {

	p := point{1, 2}
	fmt.Printf("%v\n", p)

	//If the value is a struct, the %+v variant will include the struct’s field names.
	fmt.Printf("%+v\n", p)

	// The %#v variant prints a Go syntax representation of the value, i.e.
	// the source code snippet that would produce that value.
	fmt.Printf("%#v\n", p)

	// print the Type of a value
	fmt.Printf("%T\n", p)

	// bools are easy, why it's a t no idea
	fmt.Printf("%t\n", true)

	// use d for base 10
	fmt.Printf("%d\n", 123)

	// b is for binary
	fmt.Printf("%b\n", 14)

	// c is for char corresponding tot he int
	fmt.Printf("%c\n", 33)

	// x is for hex
	fmt.Printf("%x\n", 456)

	// basic decimal formatting for floats
	fmt.Printf("%f\n", 78.9)

	// Two versions of scientific notation
	fmt.Printf("%e\n", 12340000.0)
	fmt.Printf("%E\n", 12340000.0)

	// basic string printing uses s
	fmt.Printf("%s\n", "\"string\"")

	// for double quoting strings
	fmt.Printf("%q\n", "\"string\"")

	// %x renders the string in base-16
	fmt.Printf("%x\n", "hex this")

	// use %p for representations of pointers
	fmt.Printf("%p\n", &p)

	// When printing numbers a lot of the stime you will want to control
	// the width and precision of the resulting figure
	// to specify the width of an integer use a number after the % in the verb
	// default is right justified and padded w/ spaces

	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// its albso possible to do the same with floats
	// although usually we will want to restrict the precision with width.precision
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// for left justification just use   -
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// controlling width of printed strings
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// Sprintf formats and returns the string
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// format + print to io.Writers other that os.St3dout using Fprintf

	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
