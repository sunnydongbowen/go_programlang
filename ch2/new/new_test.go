package new

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      new_test.go
// @author:    bowen
// @time:      2023-02-09 9:12
// @description:

func TestNew1(t *testing.T) {
	p := new(int)
	fmt.Println(*p) //0
	*p = 2
	fmt.Println(*p)        //2
	fmt.Printf("%T\n", p)  // *int
	fmt.Printf("%T\n", *p) // int
}

func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

func TestNew2(t *testing.T) {
	fmt.Println(newInt1())
	fmt.Println(newInt2())
}

func TestNew3(t *testing.T) {
	p := new(int)
	q := new(int)
	fmt.Println(p == q) //
}
