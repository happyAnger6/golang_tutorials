package main

func reverse(s []byte) []byte{
	var indexs []int
	j := 0
	for i, c := range string(s) {
		indexs[j]  = i
		j++
	}
}
