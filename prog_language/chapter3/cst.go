package main

import "fmt"

func main() {
	const (
		a = 1
		b
		c = 3
		d
	)
	fmt.Println(a,b,c,d)
}
