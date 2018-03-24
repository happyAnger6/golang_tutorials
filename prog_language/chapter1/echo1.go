package main

import (
	"os"
	"fmt"
)

func main() {
	var sep, s string
	for _, arg := range os.Args[1:len(os.Args)] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
