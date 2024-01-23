package structs

import "math"

// Go interface resolution is implicit.
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// it is a convention in Go to have the receiver variable
// be the first letter of the type.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Height * t.Base) * 0.5
}

// go doesn't allowed function overloading
// func Area(circle Circle) float64       {}
// func Area(rectangle Rectangle) float64 {}

// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }
