package main

import "fmt"

func maxSub(a []int) (maxStart, maxEnd, maxSub int) {
	maxSub, maxTmp := -1000000000, 0
	maxStart, maxEnd, newStart, newEnd := 0, 0, 0, 0

	for i, e := range a {
		if maxTmp <= 0 { //begin or should new start
			maxTmp = e
			newStart = i
			newEnd = i
		} else { //try add may be can be max
			maxTmp += e
			newEnd = i
		}

		if maxTmp > maxSub {
			maxSub = maxTmp
			maxStart = newStart
			maxEnd = newEnd
		}
	}

	return
}

func main() {
	a := []int{-9,0,2,4,6,-6,-2,8,8,10,-8,20,-9,-4,-20,9,2,-1,3}
	fmt.Println(maxSub(a))
}
