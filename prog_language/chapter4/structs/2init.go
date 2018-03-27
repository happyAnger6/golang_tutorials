package main

import (
	"fmt"
)

type Point struct{
	X, Y int
}

func modify(m map[string]int) {
	m["b"] = 10
}

func main() {
	p := new(Point)
	m := map[string]int{}

	m["a"] = 3
	fmt.Printf("%T %T\n", p, m)
	fmt.Println(m)
	modify(m)
	fmt.Println(m)
}
