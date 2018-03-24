package main

import "fmt"

var x uint8 = (1<<1) | (1<<5)
var y uint8 = (1<<1) | (1<<2)

func main() {
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	for i := uint(0); i < 8; i++ {
		if (x&(1<<i) != 0) {
			fmt.Println(i)
		}
	}
}

