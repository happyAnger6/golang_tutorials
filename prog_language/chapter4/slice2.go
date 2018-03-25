package main

import "fmt"

func main() {
	s := []int{}
	var s1 []int
	s2 := make([]int, 0)
	s3 := make([]int, 8, 20)
	fmt.Println(s==nil, s1==nil, s2==nil)
	fmt.Printf("s=%v s1=%v s2=%v\n", s, s1, s2)
	s3[8] = 9
	fmt.Println(s3)
}
