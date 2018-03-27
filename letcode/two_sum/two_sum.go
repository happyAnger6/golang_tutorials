package main

import "fmt"

func main() {
	nums := []int{1, 4, 6, 2, 3, 8, 10, 22, 24}
	target := 25

	rMap := map[int]int{}
	for i, num := range nums {
		r := target - num
		if j, ok := rMap[num]; ok{
			fmt.Println(i,j)
		} else {
			rMap[r] = i
		}
	}
}
