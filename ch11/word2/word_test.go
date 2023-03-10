package word

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      word_test.go
// @author:    bowen
// @time:      2023-02-12 13:14
// @description:

func TestIsPalindrome(t *testing.T) {
	// 这个就是普通的结构体定义啊！稍微复杂一点就不知道了？
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			fmt.Errorf(`IsPalindrome(%q) = %v`, test.input, got)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func BenchmarkIsPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome2("A man, a plan, a canal: Panama")
	}
}

func BenchmarkIsPalindrome3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome3("A man, a plan, a canal: Panama")
	}
}
