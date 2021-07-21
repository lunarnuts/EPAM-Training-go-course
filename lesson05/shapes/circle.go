package shapes

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() (float64, bool) {
	if _, ok := c.GetRadius(); !ok {
		fmt.Println(c)
		return -1, ok
	}
	return math.Pi * math.Pow(c.radius, 2), true
}

func (c Circle) Perimeter() (float64, bool) {
	if _, ok := c.GetRadius(); !ok {
		fmt.Println(c)
		return -1, ok
	}
	return 2 * math.Pi * c.radius, true
}

func (c Circle) String() string {
	if c.radius <= 0 {
		return fmt.Sprintf("radius needs to be positive and non-zero, got: %.3f", c.radius)
	}
	return fmt.Sprintf("\nCircle: radius %.2f", c.radius)
}

func (c Circle) GetRadius() (float64, bool) {
	if c.radius <= 0 {
		fmt.Println(c)
		return -1, false
	}
	return c.radius, true
}

func (c *Circle) SetRadius(r float64) bool {
	if r <= 0 {
		fmt.Printf("radius needs to be positive and non-zero, got: %.3f", c.radius)
		return false
	}
	c.radius = r
	return true
}
