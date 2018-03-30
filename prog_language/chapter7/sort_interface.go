package main

import (
	"sort"
	"fmt"
)

type StringSlice []string

func (p StringSlice) Len() int { return len(p) }
func (p StringSlice) Less(i, j int)  bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	names := []string{"b", "a", "c", "d", "a", "e"}
	sort.Sort(StringSlice(names))
	fmt.Println(names)
}
