package main

import "fmt"

type T3 struct {

}
type T1 struct {
	T3
}

type T2 struct {

}

func (t T3) TF2 () {
	fmt.Println("T3 T2")
}

func (t T1) TF() {
	fmt.Println("T1 T")
}

func (t T2) TF2() {
	fmt.Println("T2 T")
}

type Point struct {
	T1
	T2
}

func main() {
	p := Point{}
	fmt.Println(p)
	p.T1.TF()
	p.TF2() //breadth first.
	p.T1.TF2()
	//p.T()
}
