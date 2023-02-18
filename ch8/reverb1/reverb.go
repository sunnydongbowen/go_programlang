package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// @program:   go-basic-exercises
// @file:      reverb.go
// @author:    bowen
// @time:      2023-02-18 10:11
// @description:server端

func echo(c net.Conn, shout string, delay time.Duration) {
	//转大写输出
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)

	//原样输出
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)

	//小写输出
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func hanleConn(c net.Conn) {
	//获取输入
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "192.168.0.103:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// 接收客户端消息
		conn, err := l.Accept()
		if err != nil {
			log.Fatal()
			continue
		}
		go hanleConn(conn)
	}
}
