package main

import "fmt"
import "math"

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2 * r.width + 2 * r.height
}

// area = πr^2
func (c circle) area() float64 {
	return math.Pi * math.Exp(c.radius)
}

// perimeter = 2πr
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rect { width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}