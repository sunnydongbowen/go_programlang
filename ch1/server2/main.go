package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// @program:   go-basic-exercises
// @file:      main.go
// @author:    bowen
// @time:      2023-02-07 12:11
// @description:

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	//fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	//panic("错误")
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}
