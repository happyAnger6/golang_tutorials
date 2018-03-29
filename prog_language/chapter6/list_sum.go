package main

import "fmt"

type ListInt struct {
	Value int
	Tail *ListInt
}


func (l *ListInt) sum() int{
	if l == nil {
		return 0
	} else {
		return l.Value + l.Tail.sum()
	}
}

func main() {
	l := ListInt{
		0,
		&ListInt{1,
								&ListInt{2,
													nil}},
	}
	fmt.Println(l.sum())
}
