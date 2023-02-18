package prefix

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      prefix_test.go
// @author:    bowen
// @time:      2023-02-09 15:37
// @description:

// 源码 判断前缀
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// 源码 判断后缀
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// 源码 包含子串测试
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func TestPre(t *testing.T) {
	s1 := "bowen123"
	s2 := "23"
	fmt.Println(HasPrefix(s1, s2))
	fmt.Println(HasSuffix(s1, s2))
	fmt.Println(Contains(s1, s2))
	//strings.HasPrefix()
	//strings.Contains()
}
