package main

import "fmt"

func f() (p int) {
	defer func() interface{} {
		p = recover().(int)
		return p
	}()
	panic(1)
}

func main() {
	fmt.Println(f())
}
