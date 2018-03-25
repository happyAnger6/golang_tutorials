package main

import "fmt"

func main(){
	a := []int{1,2,3}
	b := append(a, a...)
	fmt.Println(a, b)
}
