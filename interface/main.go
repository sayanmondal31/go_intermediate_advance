package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

type Rect struct {
	Width, Height float64
}

type Circle struct {
	radius float64
}

func (r *Rect) area() float64 {
	return r.Height * r.Width
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r *Rect) perim() float64 {
	return 2 * (r.Width + r.Height)
}

func (c *Circle) perim() float64 {
	return math.Pi * c.radius * 2
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := &Rect{Width: 3, Height: 4}
	c := &Circle{radius: 5}

	measure(r)
	measure(c)

	myPrinter(1, "dsfsd", 45.6)
	value(1)

}

func value(i any) {
	switch i.(type) {
	case int:
		fmt.Println("Type: Int")
	case string:
		fmt.Println("Type: string")
	default:
		fmt.Println("Type: unknown")

	}

}

func myPrinter(i ...any) { // any is alias of interface{} it accepts any type of value
	for _, v := range i {
		fmt.Println(v)
	}
}
