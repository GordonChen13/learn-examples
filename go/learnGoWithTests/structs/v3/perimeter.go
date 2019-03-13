package v1

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.width + rect.height)
}

func (rect Rectangle) Area() float64 {
	return rect.width * rect.height
}

func (circle Circle) Perimeter() float64 {
	return 2 * circle.Radius * math.Pi
}

func (circle Circle) Area() float64 {
	return circle.Radius * circle.Radius * math.Pi
}
