package shapes

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() (float64, error) {
	if _, ok := c.GetRadius(); ok != nil {
		return -1, ok
	}
	return math.Pi * math.Pow(c.radius, 2), nil
}

func (c Circle) Perimeter() (float64, error) {
	if _, ok := c.GetRadius(); ok != nil {
		return -1, ok
	}
	return 2 * math.Pi * c.radius, nil
}

func (c Circle) String() string {
	if _, ok := c.GetRadius(); ok != nil {
		return ok.Error()
	}
	return fmt.Sprintf("\nCircle: radius %.2f", c.radius)
}

func (c Circle) GetRadius() (float64, error) {
	if c.radius <= 0 {
		return -1, fmt.Errorf("radius needs to be positive and non-zero, got: %.3f", c.radius)
	}
	return c.radius, nil
}

func (c *Circle) SetRadius(r float64) error {
	if _, ok := c.GetRadius(); ok != nil {
		return ok
	}
	c.radius = r
	return nil
}
