package main

import "fmt"

type entry struct {
	name string
	age int
}

func main() {
	var m = make(map[string]*entry)
	entries := []entry{
		{"zhangxa", 10},
		{name:"zhangxp", age:10},
	}

	for _, e := range entries {
		m[e.name] = &e
	}

	fmt.Println(m)
}
