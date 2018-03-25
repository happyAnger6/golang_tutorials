package main

import "fmt"

func appendInt(a []int, y int) []int {
	var z []int
	zlen := len(a) + 1
	if zlen <= cap(a) {
		z = a[:zlen]
	} else {
		zcap := zlen
		if zcap < 2 * len(a) {
			zcap = 2 * len(a)
		}
		z = make([]int, zlen, zcap)
		copy(z, a)
	}

	z[len(a)] = y
	return z
}

func main() {
	a := []int{1,2,3}
	fmt.Println(a, cap(a), len(a))
	b := appendInt(a, 10)
	fmt.Println(a, cap(a), len(a))
	fmt.Println(b, cap(b), len(b))

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Println(y, len(y), cap(y))
		x = y
	}
}
