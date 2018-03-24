package main

import "fmt"

var global *int

func f() {
	var l int;
	l = 1;
	global = &l
}

func main() {
	fmt.Printf("global:%v\n", global)
	f()
	fmt.Printf("global:%v\n", global)
}
