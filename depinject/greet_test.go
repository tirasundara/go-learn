package depinject

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Tira")

	got := buffer.String()
	want := "Hello, Tira"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
