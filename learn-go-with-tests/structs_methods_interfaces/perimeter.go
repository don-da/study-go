package structs_methods_interfaces

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	perimeter := (rectangle.Width + rectangle.Height) * 2
	return perimeter
}

func (r Rectangle) Area() float64 {
	area := r.Width * r.Height
	return area
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
