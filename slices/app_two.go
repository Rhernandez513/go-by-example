package main

import (
	"regexp"
	"io/ioutil"
	"fmt"
)

func main() {
	file := "dummy_data.txt"

	digits := CopyDigits(file)

	fmt.Println(digits)
}

func CopyDigits(filename string) []byte {
	var digitRegexp = regexp.MustCompile("[0-9]+")
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([] byte, len(b))
	copy(c, b)
	return c
}

// EOF
