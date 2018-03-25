package main

import "fmt"

func noempty(s []string) []string {
	i := 0
	for _, c := range s {
		if c != "" {
			s[i] = c
			i++
		}
	}
	return s[:i]
}

func main() {
	data := []string{"one", "", "three"}
	data1 := noempty(data)
	fmt.Println(data)
	fmt.Println(data1)
}
