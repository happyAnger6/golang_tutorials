package main

import "fmt"

type node struct {
	value int
	left, right *node
}

func appendValues(a []int, t * node) []int {
	if t != nil {
		a = appendValues(a, t.left)
		a = append(a, t.value)
		a = appendValues(a, t.right)
	}
	return a
}

func Sort(a []int) {
	var root *node
	for _, e := range a {
		root = add(root, e)
	}
	appendValues(a[:0], root)
}

func add(t *node, v int) *node {
	if t == nil {
		n := node{value: v}
		t = &n
		return t
	} else {
		if v < t.value {
			t.left = add(t.left, v)
		} else {
			t.right = add(t.right, v)
		}
	}
	return t
}

func main() {
	a := []int{1,6,3,2,8,10,100,4,8,7,3,111}
	Sort(a)
	fmt.Println(a)
}
