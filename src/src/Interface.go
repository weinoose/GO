package src

import (
	"fmt"
	"math"
)

type Volume interface {
	volume() float64
}

type Cube struct {
	Lenght int
}

func (c Cube) volume() float64 {
	return float64(math.Pow(float64(c.Lenght), 3))
}

func getCubeVolume(v Volume) float64 {
	return v.volume()
}

func Interface() {
	cube := Cube{Lenght: 30}
	fmt.Printf("Volume of the cube is: %f\n", getCubeVolume(cube))
}
