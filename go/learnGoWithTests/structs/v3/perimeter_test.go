package v1

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("rectangle perimeter", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		got := rect.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circle perimeter", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Perimeter()
		want := 2 * 10.0 * math.Pi

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}

	})
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangle area", func(t *testing.T) {
		rect := Rectangle{20.0, 10.0}
		want := 200.0

		checkArea(t, rect, want)
	})

	t.Run("circle area", func(t *testing.T) {
		circle := Circle{10.0}
		want := 10.0 * 10.0 * math.Pi

		checkArea(t, circle, want)
	})
}
