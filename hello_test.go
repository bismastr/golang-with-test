package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello World' when empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}
