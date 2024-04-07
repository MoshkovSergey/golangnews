package main

import (
	"fmt"
	"net/http"
	"time"
)


// main is the entry point of the Go program.
// It sets up the HTTP server to handle requests to the "/" and "/contact" routes.
func main() {
	// Print a welcome message.
	fmt.Println("I want to stand as a Go developer!")

	// Set up the routes for the HTTP server.
	// The path "/" is handled by the "pathHandler" function,
	// which serves different pages based on the requested URL path.
	http.HandleFunc("/", pathHandler)

	// Print a message indicating that the server is running on port 8080.
	fmt.Println("Server running on port 8080")

	// Start the HTTP server and listen for incoming requests on port 8080.
	// The third argument is nil, which means that the server will use the default ServeMux.
	fmt.Println("Listening for incoming requests...")
	http.ListenAndServe(":8080", nil)
}

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
	fmt.Fprintf(w, `<h1>Hello World! %s</h1>
	<form method="post" action="/contact">
	<button>Contacts</button>
	</form>`, time.Now())
}

// contactHandler is an HTTP handler function that serves the contact page.
// It sets the "Content-Type" header to "text/html; charset=utf-8"
// and writes an HTML response containing information about the developer's contacts.
// The response includes a link to go back to the home page and a link to send an email.
func contactHandler(w http.ResponseWriter, r *http.Request) {
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
	<a href="/">Назад</a>
	<a href="mailto:5Lj9Z@example.com">Написать</a><br>`)

	fmt.Fprint(w, r.URL)
}

// pathHandler is an HTTP handler function that serves different pages
// based on the requested URL path.
// It sets the "Content-Type" header to "text/html; charset=utf-8"
// and writes an HTML response containing the home page or the contact page.
// If the requested URL path is not "/" or "/contact", it sends a "404 Not Found" response.
func pathHandler(w http.ResponseWriter, r *http.Request) {
	// Switch statement checks the requested URL path and calls
	// the appropriate handler function.
	switch r.URL.Path {
	case "/": // If the path is "/", call the homeHandler.
		homeHandler(w, r)

	case "/contact": // If the path is "/contact", call the contactHandler.
		contactHandler(w, r)

	default: // If the path is anything else, send a "404 Not Found" response.
		http.NotFound(w, r)
		// TODO: add 404 page
	}
}
