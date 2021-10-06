package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		given := []int{1, 2, 3, 4, 5}
		got := Sum(given)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, given)
		}
	})

	t.Run("collection of any numbers", func(t *testing.T) {
		given := []int{1, 2, 3}
		got := Sum(given)
		want := 6

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, given)
		}
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
