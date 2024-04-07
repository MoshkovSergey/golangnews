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
	fmt.Fprintf(w, "<h1>Hello World! %s</h1><form action='/contact'><button >Контакты</button></form>", time.Now())
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Контакты</h1><p>Телефон: 8-999-999-99-99</p><p>Адрес: Москва</p><a href='/'>Назад</a><p>Время работы: 9:00-18:00</p><p>Часы работы: Пн-Пт</p>")
}
