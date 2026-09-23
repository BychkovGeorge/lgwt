package structs

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	base, height float64
}

type Rectangle struct {
	w, h float64
}

type Circle struct {
	r float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.w + r.h)
}

func (r Rectangle) Area() float64 {
	return r.w * r.h
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (t Triangle) Area() float64 {
	return t.height * t.base / 2
}
