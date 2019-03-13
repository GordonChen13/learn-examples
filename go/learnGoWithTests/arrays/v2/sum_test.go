package v1

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T)  {
	numbers := []int{1, 2, 3, 4}
	want := 10
	got := Sum(numbers)

	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T)  {
	numbers1 := []int{1, 2, 3, 4}
	numbers2 := []int{2, 3}
	want := []int{10, 5}
	got := SumAll(numbers1, numbers2)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v, given %v, %v", got, want, numbers1, numbers2)
	}
}

func TestSumAllTail(t *testing.T)  {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTail([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("make the sums of empty slices", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{4, 5, 6})
		want := []int{0, 11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
