package main

import (
	"fmt"
	"go-basic-exercises/programLanguage/ch2/tempconv"
	"os"
	"strconv"
)

// @program:   go-basic-exercises
// @file:      main.go
// @author:    bowen
// @time:      2023-02-08 15:41
// @description:

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf:%v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t) //参数类型转换
		c := tempconv.Celsius(t)
		fmt.Printf("%s=%s,%s=%s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
