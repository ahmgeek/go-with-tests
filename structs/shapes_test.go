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

	areaTest := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10.0}, 314.1592653589793},
		{Sphere{10.0}, 1256.63706143591730324260},
	}


	for _, tt := range areaTest {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
			
		}
	}
}
