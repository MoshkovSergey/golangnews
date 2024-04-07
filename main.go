package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("I want stand a golang developer!")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Server runing on port 8080")
	http.ListenAndServe(":8080", nil)

}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World! %s</h1>", time.Now())
}