package main

import (
	"fmt"
	"sort"
)

func main() {
	var names []string
	ages := map[string]int {
		"alice": 20,
		"zed": 30,
		"bob": 31,
		"an": 22,
	}

	for name, _ := range ages {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names{
		fmt.Println(ages[name])
	}
}
