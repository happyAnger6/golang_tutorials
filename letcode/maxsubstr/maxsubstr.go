package  main

import "fmt"

func maxSubstr(s string) (maxStart, maxEnd, maxLen int) {
	mapRunes := make(map[rune]int)
	start := 0
	for i, c := range s {
		if (mapRunes[c] == 0 || start > mapRunes[c]) {
			if maxLen < (i - start + 1) {
				maxLen = (i - start + 1)
				maxStart = start
				maxEnd = i
			}
		} else {
			start = mapRunes[c]
		}
		mapRunes[c] = i + 1
	}

	return
}

func main() {
	fmt.Println(maxSubstr("aabcabcbb"))
	fmt.Println(maxSubstr("aabcdabcfebb"))
}
