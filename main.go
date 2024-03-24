package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("I stand a golang developer!")
	http.HandleFunc("/", handleFunc)
	fmt.Println("Server runing on port 8080")
	http.ListenAndServe(":8080", nil)
	
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}