package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// @program:   go-basic-exercises
// @file:      echo.go
// @author:    bowen
// @time:      2023-02-12 15:56
// @description:

var (
	//默认是不换行
	//默认分割符是空格
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "separator")
)

var out io.Writer = os.Stdout

func main() {
	flag.Parse()
	if err := echo(!*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo:%v\n", err)
		os.Exit(1)
	}
}

// 这个函数其实很简单，用join函数连接切片，
func echo(newline bool, sep string, args []string) error {
	// newline 是不是新行，新行就换行，
	// seq 分隔符，
	// args 字符串切片
	fmt.Fprint(out, strings.Join(args, sep))
	if newline {
		fmt.Fprintln(out)
	}
	return nil
}
