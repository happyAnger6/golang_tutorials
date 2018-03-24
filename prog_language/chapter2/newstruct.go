package main

import "fmt"

func main() {
	fmt.Println(new(struct{}) == new(struct{}))
	fmt.Println(new([0]int) == new([0]int))
	fmt.Println(new([]int) == new([]int))
}
