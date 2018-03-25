package main

import "fmt"

func main() {
	var a [3]int;
	fmt.Println(a[0])
	fmt.Println(a[len(a) - 1])

	q := [...]int{1,2,3}
	fmt.Println(len(q))
}

