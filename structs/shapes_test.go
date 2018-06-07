package shape

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if want != got {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	var assertCorrectArea = func(t *testing.T, want float64, got float64) {
		t.Helper()
		if want != got {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}


	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := Area(rectangle)
		want := 72.0

		assertCorrectArea(t, want, got)
	})
}
