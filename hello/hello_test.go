package hello

import "testing"

func TestHello(t *testing.T) {
	assertCorrentMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Tira", "")
		want := "Hello, Tira"

		assertCorrentMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Tira", "Spanish")
		want := "Hola, Tira"

		assertCorrentMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Tira", "French")
		want := "Bonjour, Tira"

		assertCorrentMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrentMessage(t, got, want)
	})
}
