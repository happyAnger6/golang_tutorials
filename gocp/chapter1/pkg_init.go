package main

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Printf("map:%v\n", m)
	fmt.Println("init.")
	info = fmt.Sprintf("OS:%s Arch:%s", runtime.GOOS, runtime.GOARCH)
}

var m = map[string]int{"a":1, "b":2}

var info string

func main() {
	fmt.Println(info)
	c := make(chan int)
	close(c)
	for e := range c {
		fmt.Println(e)
	}
}
