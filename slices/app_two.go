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

	fmt.Println(ConciseCopyDigits(file))
}

func CopyDigits(filename string) []byte {
	var digitRegexp = regexp.MustCompile("[0-9]+")
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([] byte, len(b))
	copy(c, b)
	return c
}

func ConciseCopyDigits(filename string) []byte {
	var digitRegexp = regexp.MustCompile("[0-9]+")
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	return append(make([] byte, 0), b...)
}

// EOF
