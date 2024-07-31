package helloworld

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(got, want, t)
	})
	t.Run("say 'Hello World' when empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(got, want, t)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodia", "Spanish")
		want := "Hola, Elodia"

		assertCorrectMessage(got, want, t)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodia", "French")
		want := "Bonjour, Elodia"

		assertCorrectMessage(got, want, t)
	})
}

func assertCorrectMessage(got, want string, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
