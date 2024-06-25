package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

// homeHandler is an HTTP handler function that serves the home page.
// It sets the "Content-Type" header to "text/html; charset=utf-8"
// and writes an HTML response containing the current time and a form
// that redirects to the "/contact" route.
func homeHandler(w http.ResponseWriter, _ *http.Request) {
	// Set the "Content-Type" header to indicate that the response
	// is an HTML document in UTF-8 encoding.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Use fmt.Fprintf to write the HTML response to the ResponseWriter.
	// The first argument is the ResponseWriter, to which the formatted
	// string is written. The second argument is a formatted string
	// that includes the current time and a form with a button that
	// redirects to the "/contact" route when clicked.
	// for {}
	fmt.Fprintf(w, `<h1>Привет мир! Сейчас в Санкт-Петербурге: %s.</h1>
	<form method="get" action="/contact">
	<button>Contacts</button>
	</form>
	<form method="get" action="/faq">
	<button>FAQ</button>
	</form>
	`, time.Now().Format(" 15:04:05 02-01-2006"))
}

// contactHandler is an HTTP handler function that serves the contact page.
// It sets the "Content-Type" header to "text/html; charset=utf-8"
// and writes an HTML response containing information about the developer's contacts.
// The response includes a link to go back to the home page and a link to send an email.
func contactHandler(w http.ResponseWriter, _ *http.Request) {
	// Set the "Content-Type" header to indicate that the response
	// is an HTML document in UTF-8 encoding.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Use fmt.Fprintf to write the HTML response to the ResponseWriter.
	// The first argument is the ResponseWriter, to which the formatted
	// string is written. The second argument is a formatted string
	// that includes information about the developer's contacts and
	// links to go back to the home page and send an email.
	fmt.Fprint(w, `<h1>Контакты</h1>
	<p>Sergey</p>
	<p>Telegram: @sergey</p>
	<a href="/">Назад</a><br>`)
}

func faqHandler(w http.ResponseWriter, _ *http.Request) {
	// Set the "Content-Type" header to indicate that the response
	// is an HTML document in UTF-8 encoding.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Use fmt.Fprintf to write the HTML response to the ResponseWriter.
	// The first argument is the ResponseWriter, to which the formatted
	// string is written. The second argument is a formatted string
	// that includes information about the developer's contacts and
	// links to go back to the home page and send an email.
	fmt.Fprint(w, `<h1>FAQ</h1>
	<p>Sergey</p>
	<p>Sergey</p>
	<p>Sergey</p>
	<p>Sergey</p>
	<a href="/">Назад</a><br>
	`)
}

// Router is a struct that satisfies the http.Handler interface.
type Router struct{}

// ServeHTTP is the method that satisfies the http.Handler interface.
// It calls the pathHandler function to serve different pages based on the requested URL path.
// It takes in a http.ResponseWriter and a *http.Request as parameters.
// It does not return anything.
func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/": // If the path is "/", call the homeHandler.
		homeHandler(w, r)

	case "contact": // If the path is "/contact", call the contactHandler.
		contactHandler(w, r)

	case "faq": // If the path is "/contact", call the contactHandler.
		faqHandler(w, r)
	default: // If the path is anything else, send a "404 Not Found" response.
		http.Error(w, "страница не найдена", http.StatusNotFound)
	}
	// Call the pathHandler function to serve different pages based on the requested URL path.
	// The pathHandler function takes in a http.ResponseWriter and a *http.Request as parameters.
	// It does not return anything.
}

// main is the entry point of the Go program.
// It sets up the HTTP server to handle requests to the "/" and "/contact" routes.
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// var router Router
	// Print a message indicating that the server is running on port 8080.
	fmt.Println("Server running on port 8080")

	// Start the HTTP server and listen for incoming requests on port 8080.
	// The third argument is nil, which means that the server will use the default ServeMux.
	fmt.Println("Listening for incoming requests...")
	r.Group(func(r chi.Router) {
		r.Get("/", homeHandler)
		r.Get("/contact", contactHandler)
		r.Get("/faq", faqHandler)
	})
	// Start the HTTP server and listen for incoming requests on port 8080.
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}
