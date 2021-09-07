package shapes

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) String() string {
	return fmt.Sprintf("\nCircle: radius %.2f", c.radius)
}

func (c Circle) GetRadius() float64 {
	return c.radius
}

func (c *Circle) SetRadius(r float64) {
	c.radius = r
}
