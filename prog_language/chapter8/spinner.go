package main

import (
	"time"
	"fmt"
)

func spinner(delay time.Duration) {
	for _, c := range `-\|/` {
		fmt.Printf("\r%c", c)
		time.Sleep(delay)
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n) + fib(n-1)
}

func fib1(n int64) int64 {
	var x, y int64 = 0, 1
	for i:= int64(1); i <= n; i++ {
		x, y = y, x+y
	}

	return x
}

func main() {
	go spinner(1000 * time.Millisecond)
	f := fib1(100000000000)
	fmt.Println()
	fmt.Println(f)
}
