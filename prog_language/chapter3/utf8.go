package main

import "fmt"

func main() {
	s := "\u4e16\u754c"
	fmt.Println("\u4e16\u754c", len(s))
	fmt.Printf("type:%T\n", s)
	for i, _ := range s {
		fmt.Printf("i:%d %x type:%[2]T", i, s[i])
	}
}
