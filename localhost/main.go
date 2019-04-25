package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", handler2)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HOMEPAGE")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO WORLD")
}
