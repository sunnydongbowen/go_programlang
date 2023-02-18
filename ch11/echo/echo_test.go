package main

import (
	"bytes"
	"fmt"
	"testing"
)

// @program:   go-basic-exercises
// @file:      echo_test.go
// @author:    bowen
// @time:      2023-02-12 20:31
// @description:

func TestEcho(t *testing.T) {
	var tests = []struct {
		newline bool
		seq     string
		args    []string

		want string
	}{
		{true, "", []string{}, "\n"},
		{false, "", []string{}, ""},
		{true, "\t", []string{"one", "two", "three"}, "one\ttwo\tthree\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
		{false, ":", []string{"1", "2", "3"}, "1:2:3"},
		{true, ",", []string{"a", "b", "c"}, "a b c\n"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%v,%q,%q)", test.newline, test.seq, test.args)
		out = new(bytes.Buffer)
		if err := echo(test.newline, test.seq, test.args); err != nil {
			t.Errorf("%s failed:%v", descr, err)
			continue
		}
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s=%q,want %q", descr, got, test.want)
		}
	}
}
