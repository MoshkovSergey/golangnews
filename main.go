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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Hello World! %s</h1><form action=\"/contact\"><button>Контакты</button></form>", time.Now())
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Контакты</h1><p>Sergey</p><p>Telegram: @sergey</p><a href=\"/\">Назад</a>")
}