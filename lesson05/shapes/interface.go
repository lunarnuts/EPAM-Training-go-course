package shapes

type Shape interface {
	Area() (float64, bool)
	Perimeter() (float64, bool)
}
