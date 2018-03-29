package main

import "fmt"

func main() {
	var x interface{} = []int{1, 2, 3}
	fmt.Printf("%T\n", x)
	//fmt.Println(x == x)
}
