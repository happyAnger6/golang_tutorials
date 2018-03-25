package main

import (
	"fmt"
)

func main() {
	s := "hello the world."
	months := [...]string{"January", "Februray", "Wednsday", "Aprial", "March", "June", "July", "Angust", "Sempteber"}
	a := months[4:]
	fmt.Printf("a:%v %[1]T\n",a)
	fmt.Println(months)
	fmt.Printf("%T\n", s[2:3])
}
