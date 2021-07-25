package main

import (
	"fmt"
	"log"

	"github.com/lunarnuts/go-course/tree/lesson05/shapes"
)

func DescribeShape(s shapes.Shape) {
	a, ok := s.Area()
	if ok != nil {
		log.Fatal(ok)
	}
	p, ok := s.Perimeter()
	if ok != nil {
		log.Fatal(ok)
	}
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", a)
	fmt.Printf("Perimeter: %.2f\n", p)
}

func main() {
	var c shapes.Circle
	c.SetRadius(8)
	var r shapes.Rectangle
	r.SetHeight(9)
	r.SetWidth(3)
	DescribeShape(c)
	DescribeShape(r)
	var a shapes.Circle
	a.SetRadius(2.5)
	DescribeShape(a)
}
