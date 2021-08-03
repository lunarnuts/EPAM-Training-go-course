package shapes

type Shape interface {
	Area() (float64, error)
	Perimeter() (float64, error)
}
