package main

type Shape struct {
	Rectangle
}

type Rectangle struct {
	Length float64
	Width  float64
}

// method with value receiver
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r *Rectangle) Scale(factor float64) {
	r.Length *= factor
	r.Width *= factor
}
