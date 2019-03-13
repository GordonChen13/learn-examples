package v1

type Rectangle struct {
	width float64
	height float64
}

func (rect Rectangle) Perimeter() float64 {
	return 2*(rect.width + rect.height)
}

func (rect Rectangle) Area() float64 {
	return rect.width * rect.height
}
