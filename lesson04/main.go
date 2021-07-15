package main

import (
	"fmt"
	"github.com/lunarnuts/go-course/tree/lesson04/shapes"
)

func DescribeShape(s shapes.Shape) {
   fmt.Println(s)
   fmt.Printf("Area: %.2f\n", s.Area())
   fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {


   var c shapes.Circle
   c.SetRadius(8)
   var r shapes.Rectangle
   r.SetHeight(9)
   r.SetWidth(3)
   DescribeShape(c)
   DescribeShape(r)
}
