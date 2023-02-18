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
// @time:      2023-02-18 11:08
// @description: reverb1 测试客户端
func main() {
	conn, err := net.Dial("tcp", "192.168.0.103:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
