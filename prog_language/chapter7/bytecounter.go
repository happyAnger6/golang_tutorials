package main

import "fmt"

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	var name = "Dollay"
	fmt.Fprintf(&c, "Hello %s", name)
	fmt.Println(c)
	fmt.Printf("%%a\n")
}
