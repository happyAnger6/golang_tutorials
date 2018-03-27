package main

import "fmt"

type Point struct{
	X, Y int
}

type point1 struct {
	Z, W int
}
type Circle struct {
	 a Point
	 Point
	 point1
	 Radius float64
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{0, 1},Point{1,2}, point1{3,4,}, 10.2}, 100}

	fmt.Println(w)
	fmt.Println(w.Circle.Point.X)
	fmt.Println(w.Circle.a.X)
	fmt.Printf("%#v\n", w)
	fmt.Println(w.point1.Z)
	fmt.Println(w.W)
}
