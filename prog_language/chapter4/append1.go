package main

import "fmt"

func main() {
	var runes []rune
	for _, s := range "hello, 世界" {
		runes = append(runes, s)
	}
	fmt.Println(runes)
}
