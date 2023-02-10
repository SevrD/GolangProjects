package main

import (
	"fmt"
	"math"
)

type Cube struct {
	side int
}

type Sphere struct {
	radius int
}

type Cylinder struct {
	radius int
	height int
}

func (c *Cube) init(side int) {
	(*c).side = side
}

func (c *Sphere) init(r int) {
	(*c).radius = r
}

func (c *Cylinder) init(r int, h int) {
	(*c).radius, (*c).height = r, h
}

func (c Cube) volume() float64 {
	return float64(c.side * c.side * c.side)
}

func (c Sphere) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(float64(c.radius), 3)
}

func (c Cylinder) volume() float64 {
	return math.Pi * math.Pow(float64(c.radius), 2) * float64(c.height)
}

func (c Cube) area() float64 {
	return float64(6 * c.side * c.side)
}

func (c Sphere) area() float64 {
	return 4 * math.Pi * math.Pow(float64(c.radius), 2)
}

func (c Cylinder) area() float64 {
	return 2 * math.Pi * float64(c.radius) * float64(c.height)
}

type Object3D interface {
	volume() float64
	area() float64
}

func main() {

	fig1 := Cube{23}
	fig2 := Sphere{14}
	fig3 := Cylinder{12, 34}

	figures := [...]Object3D{fig1, fig2, fig3}

	for _, fig := range figures {
		fmt.Println("Volume:", fig.volume())
		fmt.Println("Area:", fig.area())
	}

}
