package strng

import "testing"

var assertCorrectMessage = func(t *testing.T, expected, result bool) {
	t.Helper()
	if result != expected {
		t.Errorf("expected '%t' got '%t'", expected, result)
	}
}

func TestContains(t *testing.T){
	t.Run("Contains substring", func(t *testing.T) {
		result := Contains("seafood", "foo")
		expected := true

		assertCorrectMessage(t, expected, result)
	})

	t.Run("Doesn't Contain substring", func(t *testing.T) {
		result := Contains("seafood", "achmed")
		expected := false

		assertCorrectMessage(t, expected, result)
	})
}

func TestContainsAny(t *testing.T){
	t.Run("Contains a char", func(t *testing.T) {
		result := ContainsAny("seafood", "s")
		expected := true

		assertCorrectMessage(t, expected, result)
	})

	t.Run("Contains multiple chars", func(t *testing.T) {
		result := ContainsAny("seafood", "d & foo & sea")
		expected := true

		assertCorrectMessage(t, expected, result)
	})
	

	t.Run("Doesn't Contain substring", func(t *testing.T) {
		result := ContainsAny("seafood", "g")
		expected := false

		assertCorrectMessage(t, expected, result)
	})
}
