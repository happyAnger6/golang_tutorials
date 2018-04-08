package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated" = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("detartrated" = false`)
	}
	if IsPalindrome("中国中") {
		t.Error(`IsPalindrome("detartrated" = true`)
	}
}
