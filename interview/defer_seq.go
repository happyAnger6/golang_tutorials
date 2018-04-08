package main

import "fmt"

func defer_call()  {
	defer func() { fmt.Println("print before.") }()
	defer func() { fmt.Println("print middle.") }()
	defer func() { fmt.Println("print after.") }()
	panic("panic")
}

func main() {
	panic("main panic")
	defer_call()
}
