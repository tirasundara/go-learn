package slices

import (
	"fmt"
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
