package main

import "strings"
import "fmt"

//Returns the first index of the target string t, or -1 if no match is found.
func Index(vs [] string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

//Returns true if the target string t is in the slice.
func Include(vs [] string, t string) bool {
	return Index(vs, t) >= 0
}

//Returns true if one of the strings in the slice satisfies the predicate f.
func Any(vs [] string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

//Returns true if all of the strings in the slice satisfy the predicate f.
func All(vs [] string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//Returns a new slice containing all strings in the slice that satisfy the predicate f.
func Filter(vs []string, f func(string) bool) []string{
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// map a func to all strs in str arr, returning new str slice w/ results of operation
func Map(vs [] string, f func(string) string) [] string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}


func main() {
	var strs = []string {"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pears"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))


	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
}