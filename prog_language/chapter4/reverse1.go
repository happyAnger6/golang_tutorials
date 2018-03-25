package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i,j = i+1, j-1 {
		s[i], s[j] = s[j] , s[i]
	}
}

func main() {
	a := [8]int{1,2,3,4,5,6,7,8}
	fmt.Println(a)

	reverse(a[:3])
	reverse(a[3:])
	reverse(a[:])
	fmt.Println(a)

}
