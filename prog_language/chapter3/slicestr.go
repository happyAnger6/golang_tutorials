package main

import "fmt"

func main() {
	s := "hello"
	b := []byte(s)
	t := string(b)
	b[1] = 'a'
	fmt.Println(s, b, t)
}
