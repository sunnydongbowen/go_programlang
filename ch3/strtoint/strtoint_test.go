package strtoint

import (
	"fmt"
	"strconv"
	"testing"
)

// @program:   go-basic-exercises
// @file:      strtoint_test.go
// @author:    bowen
// @time:      2023-02-09 17:06
// @description:

func TestIntToStr(t *testing.T) {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

}

func TestStrToInt(t *testing.T) {
	x, _ := strconv.Atoi("123")
	y, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(x, y)

}
