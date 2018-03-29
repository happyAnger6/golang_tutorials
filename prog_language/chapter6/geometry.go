package main

import (
	"math"
	"fmt"
	"bytes"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(o Point) float64 {
	return math.Hypot(o.X - p.X, o.Y - p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type Path []Point

func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}

	return sum
}

func main() {
	fmt.Println(Point{3.9, 4.2}.Distance(Point{2,3}))
	perim := Path{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 4},
	}
	fmt.Println(perim.Distance())

	p := Point{3,4}
	p.ScaleBy(10)
	fmt.Println(p)

	//Point{1,2}.ScaleBy(20)
}
