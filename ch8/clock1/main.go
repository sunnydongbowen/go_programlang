package main

import (
	"io"
	"log"
	"net"
	"time"
)

// @program:   go-basic-exercises
// @file:      main.go
// @author:    bowen
// @time:      2023-02-17 14:18
// @description:

func main() {
	//监听端口
	listener, err := net.Listen("tcp", "192.168.0.103:8000")
	if err != nil {
		log.Fatal(err)
	}
	// 不断去接收客户端的请求
	for {
		// 一个链接一个conn
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("2006-01-02 15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
