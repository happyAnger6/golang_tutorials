package main

import "fmt"

func main() {
	var c1 chan int
	var c2 chan int
	select {
	case c1 <- 0:
		fmt.Println("case c1.")
	case c2 <- 0:
		fmt.Println("case c1.")
	default:
		fmt.Println("default.")
	}
}
