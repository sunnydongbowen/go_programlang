package main

import (
	"log"
	"net"
)

// @program:   go-basic-exercises
// @file:      main.go
// @author:    bowen
// @time:      2023-02-17 12:18
// @description:

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {

	}()
}
