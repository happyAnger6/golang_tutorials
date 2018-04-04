package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown. Press return to abort.")
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	select {
	case <-time.After(10 * time.Second):
		fmt.Printf("count.\r\n")
	case <-abort:
		fmt.Println("abort.")
	}
}
