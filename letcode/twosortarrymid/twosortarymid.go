package main

import "fmt"

func mid(a, b []int) (int) {
	sum, num := 0, 0
	i, j, tmp, mid, n := 0, 0, 0, 0, 0
	m := len(a) + len(b)
	if m%2 == 0 {
		mid = m/2
		num = 2
	} else {
		mid = m/2 + 1
		num = 1
	}

	for mid > 0 && (i < len(a) || j < len(b)) {
		n++
		if (i >= len(a)) {
			tmp = a[j]
			j++
		} else if(j >= len(b)) {
			tmp = a[i]
			i++
		} else if(a[i] <= b[j]) {
			tmp = a[i]
			i++
		} else {
			tmp = b[j]
			j++
		}

		if(n == mid) {
			if(num == 1) {
				sum+=tmp
				break;
			}
			if(num == 2) {
				sum+=tmp
				num--
				mid = mid+1
			}
		}
	}

	return sum
}

func main() {
	fmt.Println(mid([]int{2,5,6,7}, []int{1}))
}