package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("c1:%x c2:%x c1==c2:%t %T\n", c1, c2, c1==c2, c1)
}
