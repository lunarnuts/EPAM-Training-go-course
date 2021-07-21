package shapes

import (
	"fmt"
)

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) Area() (float64, bool) {
	h, ok := r.GetHeight()
	if !ok {
		return -1, ok
	}
	w, ok := r.GetWidth()
	if !ok {
		return -1, ok
	}
	return h * w, true
}

func (r Rectangle) Perimeter() (float64, bool) {
	h, ok := r.GetHeight()
	if !ok {
		return -1, ok
	}
	w, ok := r.GetWidth()
	if !ok {
		return -1, ok
	}
	return 2 * (h + w), true
}

func (r Rectangle) String() string {
	h, ok := r.GetHeight()
	if !ok {
		return fmt.Sprintf("height needs to be positive and non-zero, got: %.3f", r.height)
	}
	w, ok := r.GetWidth()
	if !ok {
		return fmt.Sprintf("width needs to be positive and non-zero, got: %.3f", r.width)
	}
	return fmt.Sprintf("\nRectangle with height %.2f and width %.2f", h, w)
}

func (r *Rectangle) SetHeight(h float64) bool {
	if h <= 0 {
		fmt.Printf("height needs to be positive and non-zero, got: %.3f\n", h)
		return false
	}
	r.height = h
	return true
}

func (r *Rectangle) SetWidth(w float64) bool {
	if w <= 0 {
		fmt.Printf("width needs to be positive and non-zero, got: %.3f\n", w)
		return false
	}
	r.width = w
	return true
}

func (r Rectangle) GetHeight() (float64, bool) {
	if r.height <= 0 {
		return -1, false
	}
	return r.height, true
}

func (r Rectangle) GetWidth() (float64, bool) {
	if r.width <= 0 {
		return -1, false
	}
	return r.width, true
}
