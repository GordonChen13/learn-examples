package v1

import "testing"

func TestSum(t *testing.T)  {
	numbers := [5]int{1, 2, 3, 4, 5}
	want := 15
	got := sum(numbers)

	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, numbers)
	}
}
