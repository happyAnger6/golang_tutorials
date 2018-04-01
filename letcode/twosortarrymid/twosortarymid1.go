package main

import "fmt"

func findMid(a, b []int) int {
	totalLen := len(a) + len(b)
	sum, n, i, j := 0, 0
	tmp, cur, next, curIndex := 0, 0, 0, 0


	for i < len(a) || j < len(b) {
		n++
		if(i > len(a)) {
			tmp = b[j]
			j++
		} else if( j > len(b)) {
			tmp = a[i]
			i++
		} else if(a[i] <= a[j]) {
			tmp = a[i]
			i++
		} else {
			tmp = b[j]
			j++
		}

		if(n==1) {
			cur = tmp
			curIndex = 1
		}
		if(n == curIndex+1) {
			next = tmp
		}

		if(n%2 == 1) {
			cur=next
			next=0
		}
	}
}

func main() {
	fmt.Println(mid([]int{2,5,6,7}, []int{1}))
}