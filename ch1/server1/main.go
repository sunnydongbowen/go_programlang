package main

import (
	"fmt"
	"log"
	"net/http"
)

// @program:   go-basic-exercises
// @file:      main.go
// @author:    bowen
// @time:      2023-02-07 11:50
// @description:

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}
