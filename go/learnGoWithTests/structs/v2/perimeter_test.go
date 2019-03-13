package v1

import "testing"

func TestPerimeter(t *testing.T)  {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T)  {
	rect := Rectangle{10.0, 10.0}
	got := rect.Area()
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
