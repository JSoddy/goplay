package main

import "fmt", "math"

type Shape interface {
	perimeter() float64
}

type Rectangle struct {
	x, y, length, height float64
}

type Circle struct {
	x, y, radius float64
}

func (c *Circle) perimeter() float64 {

	perim := c.radius * (math.Pi * 2)
	return perim

}

func (r *Rectangle) perimeter() float64 {

	perim := r.length * 2 + r.height * 2
	return perim

}
