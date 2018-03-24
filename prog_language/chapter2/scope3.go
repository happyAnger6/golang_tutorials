package main

import "fmt"

func f1() int {
	return 1
}

func g(x int) int {
	return x+1
}

func main(){
	if x := f1(); x == 0 {
		fmt.Println("x == 0")
	} else if y := g(x); x == y {
		fmt.Println("x == y")
	} else {
		fmt.Println(x,y)
	}

//	fmt.Println(x, y)
}
