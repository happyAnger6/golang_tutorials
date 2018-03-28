package main

import (
	"fmt"
	"runtime"
	"os"
)

func f(x int) {
	defer fmt.Printf("defer %d\n", x)
	fmt.Printf("f(%d) \n", x + 0/x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func main() {
	defer printStack()
	f(3)
}
