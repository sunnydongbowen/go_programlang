package tempconv

import (
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      convert_test.go
// @author:    bowen
// @time:      2023-02-08 10:33
// @description:

func TestOne(t *testing.T) {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	BoilingF := CToF(BoilingC)
	fmt.Printf("%g\n", BoilingF)
	fmt.Printf("%g\n", BoilingF-CToF(FreezingC))
	//fmt.Printf("%g\n", BoilingF-FreezingC)) 报错，类型不一致
}

func TestTwo(t *testing.T) {
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	//fmt.Println(c==f)
	fmt.Println(c == Celsius(f))
}

func TestPa(t *testing.T) {
	fmt.Printf("Brrr! %v\n", AbsoluteZeroC)
	fmt.Println(CToF(BoilingC))
}

func TestFour(t *testing.T) {
	c := FToC(212.0)
	fmt.Println(c.String())

	fmt.Printf("%v\n", c)
	//fmt.Printf("%s\v", c)
	fmt.Println(c)
	fmt.Printf("%g %T\n", c, c)
	fmt.Println(float64(c))
}
