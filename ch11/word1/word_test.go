package word

import "testing"

// @program:   go-basic-exercises
// @file:      word_test.go
// @author:    bowen
// @time:      2023-02-11 18:03
// @description:

func TestIsPalindrome(t *testing.T) {
	// Is返回true,!就是false，所以不会抛错
	// Is返回false，!就是ok
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	//
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	// 返回false,所以不会打印
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	// Is返回true !就是false
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}
