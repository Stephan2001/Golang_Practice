package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	raduis float64
}

func (c Circle) Area() float64 {
	return c.raduis * c.raduis * math.Pi
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) * 0.5
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}
