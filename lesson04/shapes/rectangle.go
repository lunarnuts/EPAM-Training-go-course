package shapes

import (
	"fmt"
)

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("\nRectangle with height %.2f and width %.2f", r.height, r.width)
}

func (r *Rectangle) SetHeight(h float64) {
	r.height = h
}

func (r *Rectangle) SetWidth(w float64) {
	r.width = w
}

func (r Rectangle) GetHeight() float64 {
	return r.height
}

func (r Rectangle) GetWidth() float64 {
	return r.width
}
