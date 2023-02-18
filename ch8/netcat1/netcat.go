package main

import (
	"io"
	"log"
	"net"
	"os"
)

// @program:   go-basic-exercises
// @file:      netcat.go
// @author:    bowen
// @time:      2023-02-17 15:56
// @description: 测试clock1 client

func main() {
	conn, err := net.Dial("tcp", "192.168.0.103:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustcopy(os.Stdout, conn)
}

func mustcopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
