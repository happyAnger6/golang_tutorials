package main

import "fmt"

type node struct {
	value int
	left, right *node
}

func add(t *node, val int) *node{
	if t == nil {
		n := node{value: val}
		t = &n
	} else {
		if val < t.value {
			t.left = add(t.left, val)
		} else {
			t.right = add(t.right, val)
		}
	}
	return t
}

func sort(ary []int) []int{
	var root *node
	for _, a := range ary {
		root = add(root, a)
	}
	return appendValues(ary[:0], root)
}

func appendValues(values []int, t *node) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func main() {
	a := []int{2,5,1,7,8,10,2,123,5,8,8,100}
	sort(a)
	fmt.Println(a)

	s := a[:3]
	n := append(s, 4)
	fmt.Println(len(s),cap(s), len(n), cap(n))
}
