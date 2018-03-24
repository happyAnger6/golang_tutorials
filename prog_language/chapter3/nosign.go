package main

import "fmt"

func main() {
	medols := []string{"gold", "silver", "bronze"}
	for i := len(medols) - 1; i >= 0; i-- {
		fmt.Printf("i:%T :%v\n", i, i)
		fmt.Println(medols[i])
	}
}
