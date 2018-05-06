package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	name := "Eric"
	go func() {
		fmt.Printf("Hello :%s\r\n", name)
	}()
	runtime.Gosched()
	name = "Harry"
	time.Sleep(time.Millisecond)
}
