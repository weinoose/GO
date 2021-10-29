package src

// STRUCTS & INTERFACES

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type calculate interface {
	calc() float64
}

func geometry(c calculate) {
	fmt.Println("Area of the shape is: ", c.calc())
}

func (s square) calc() float64 {
	return s.length * s.length
}

func (c circle) calc() float64 {
	return math.Pi * c.radius * c.radius
}

func OOP() {
	s := square{length: 17}
	c := circle{radius: 12}
	geometry(s)
	geometry(c)
}
