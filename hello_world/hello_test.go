package main

import "testing"

func assertHelper(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Drew", "")
		want := "Hello, Drew"
		assertHelper(t, got, want)
	})

	t.Run("say 'Hello, World' when there is not a string provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertHelper(t, got, want)
	})

	t.Run("Run in spanish", func(t *testing.T) {
		got := Hello("Drew", "Spanish")
		want := "Hola, Drew"
		assertHelper(t, got, want)
	})
}
