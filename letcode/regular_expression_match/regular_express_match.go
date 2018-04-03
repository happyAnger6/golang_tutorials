package main

import "fmt"

func isMatch(s, p string) bool {
	if len(p) == 0 {
		if len(s) == 0 {
			return true
		}
		return false
	}

	if len(p) == 1 {
		if len(s) == 1 && (s[0] == p[0] || p[0] == '.') {
			return true
		}
		return false
	}

	if p[1] != '*' {
		if s[0] == p[0] || p[0] == '.' {
			return isMatch(s[1:], p[1:])
		}
		return false
	}


	 for len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
		if isMatch(s[1:], p[2:]) { //最短匹配
			return true
		}
		s = s[1:] //最长匹配
	}

	return isMatch(s, p[2:])

}

func main() {
	fmt.Println(isMatch("aabcd", "a.*cd"))
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "aa"))
	fmt.Println(isMatch("aaa", "aa"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aa", ".*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b*"))
}
