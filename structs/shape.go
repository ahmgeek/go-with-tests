package shape

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Sphere struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}


func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (s Sphere) Area() float64 {
	// Amazing math skills.
	return (math.Pi * 4) * (s.Radius * s.Radius)
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
