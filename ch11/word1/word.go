package word

// @program:   go-basic-exercises
// @file:      word.go
// @author:    bowen
// @time:      2023-02-11 17:56
// @description:

func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
