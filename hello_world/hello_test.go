package main

import "testing"

func TestHello(t *testing.T){
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Ahmad!", "")
		want := "Hello, Ahmad!"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say hello world when an empty String supplied", func(t *testing.T) {

		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {

		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})
	
	t.Run("in french", func(t *testing.T) {

		got := Hello("Dodo", "French")
		want := "Bonjour, Dodo"

		assertCorrectMessage(t, got, want)
	})
}
