package stringmethods

import "strings"

func IsPalindrome(s string) bool {
	lowerRune := []rune(strings.ReplaceAll(strings.ToLower(s), " ", ""))

	for i := 0; i < len(lowerRune)/2; i++ {
		if lowerRune[i] != lowerRune[len(lowerRune)-i-1] {
			return false
		}
	}

	return true
}
