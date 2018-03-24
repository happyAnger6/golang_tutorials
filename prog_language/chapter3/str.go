package main

import "fmt"

func main() {
	s := "Hello world!"
	t := s

	fmt.Println(len(s))

	for i, c := range s {
		fmt.Printf("i:%d type:%T %T\n", i, c, s[i])
	}

	s += ", right foot."
	fmt.Println(s)
	fmt.Println(t)
}
