package main

import "fmt"

func main() {
	f := 1e100
	i := int(f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%d\n", i)
}
