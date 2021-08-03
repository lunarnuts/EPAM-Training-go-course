package shapes

import (
	"fmt"
)

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) Area() (float64, error) {
	h, ok := r.GetHeight()
	if ok != nil {
		return -1, ok
	}
	w, ok := r.GetWidth()
	if ok != nil {
		return -1, ok
	}
	return h * w, nil
}

func (r Rectangle) Perimeter() (float64, error) {
	h, ok := r.GetHeight()
	if ok != nil {
		return -1, ok
	}
	w, ok := r.GetWidth()
	if ok != nil {
		return -1, ok
	}
	return 2 * (h + w), nil
}

func (r Rectangle) String() string {
	h, ok := r.GetHeight()
	if ok != nil {
		return ok.Error()
	}
	w, ok := r.GetWidth()
	if ok != nil {
		return ok.Error()
	}
	return fmt.Sprintf("\nRectangle with height %.2f and width %.2f", h, w)
}

func (r *Rectangle) SetHeight(h float64) error {
	if _, ok := r.GetHeight(); ok != nil {
		return ok
	}
	if h <= 0 {
		return fmt.Errorf("height needs to be positive and non-zero, got: %.3f", r.height)
	}
	r.height = h
	return nil
}

func (r *Rectangle) SetWidth(w float64) error {
	if _, ok := r.GetWidth(); ok != nil {
		return ok
	}
	if w <= 0 {
		return fmt.Errorf("width needs to be positive and non-zero, got: %.3f", r.width)
	}
	r.width = w
	return nil
}

func (r Rectangle) GetHeight() (float64, error) {
	if r.height <= 0 {
		return -1, fmt.Errorf("height needs to be positive and non-zero, got: %.3f", r.height)
	}
	return r.height, nil
}

func (r Rectangle) GetWidth() (float64, error) {
	if r.width <= 0 {
		return -1, fmt.Errorf("width needs to be positive and non-zero, got: %.3f", r.width)
	}
	return r.width, nil
}
