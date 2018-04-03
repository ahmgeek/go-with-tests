package main

import "testing"

func TestHello(t *testing.T){
	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Ahmad!")
		want := "Hello, Ahmad!"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("say hello world when an empty String supplied", func(t *testing.T) {

		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
