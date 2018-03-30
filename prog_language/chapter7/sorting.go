package main

import (
	"time"
	"sort"
	"fmt"
)

type Track struct  {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var tracks = []*Track {
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m38s")},
	{"Go Ahead", "Alicia Keys", "As I am", 2007, length("3m38s")},
	{"Reday 2 Go", "Martin Solveig", "Smash",2011, length("3m38s")},
}

func main() {
	names := []int{2,3,1}
	sort.Reverse(names)
	fmt.Println(names)
}
