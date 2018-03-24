package main

import "fmt"

func main() {
	s := "你好中国"
	fmt.Println(len(s))

	fmt.Printf("% x\n", s)

	r := []rune(s)
	fmt.Println(len(r))

	for i, _ := range r {
		fmt.Printf("%d %c %[2]T\n", i, r[i])
	}
}
