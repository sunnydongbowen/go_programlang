package main

import (
	"io"
	"log"
	"net"
	"time"
)

// @program:   go-basic-exercises
// @file:      clock.go
// @author:    bowen
// @time:      2023-02-17 18:49
// @description:

func handleconn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("2006-01-02 15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "192.168.0.103:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleconn(conn)
	}

}
