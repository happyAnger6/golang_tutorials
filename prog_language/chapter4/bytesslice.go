package main

import (
	"fmt"
	"bytes"
)

func main() {
	a := []byte{0x01, 0x02}
	b := []byte{0x01, 0x02}
	fmt.Println(bytes.Equal(a, b))
}
