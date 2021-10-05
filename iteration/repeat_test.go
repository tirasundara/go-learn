package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("returns repeated string correctly", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	})

	t.Run("returns empty string when n is less than 1", func(t *testing.T) {
		got := Repeat("a", -1)
		want := ""

		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
