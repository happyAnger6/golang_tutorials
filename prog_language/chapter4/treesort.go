package main

type node struct {
	value int
	left, right *node
}

func Sort(a []int) {
	root := node{}
	for _, e := range a {
		add(root, e)
	}
}

func add(t *node, v int) {
	if t == nil {

	}
}
