package main

import "fmt"

func main() {
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", RMB: "&"}
	fmt.Println(symbol, len(symbol))
	fmt.Println(USD, EUR, GBP, RMB)

	r := [...]int{99:-1}
	fmt.Printf("%v %[1]T\n",r)

	a := [...]int{1,2}
	b := [2]int{1,2}
	c := [...]int{1,3}
	fmt.Println(a==b, b==c, a==c)
}
