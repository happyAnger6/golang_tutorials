package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello 世界"
	fmt.Println(len(s))
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c size:%d\n", i, r, size)
		i += size
	}

	for i, c := range s {
		fmt.Printf("i=%d c=%T c=%[2]c\n", i, c)
	}

	for i := len(s) - 1; i >= 0; i-- {
		fmt.Printf("i:%d c:%c\n", i, s[i])
	}
}
