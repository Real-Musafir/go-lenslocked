package main

import (
	"fmt"
	"net/http"
)

func pageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Sorry page not found</h1>")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my first page in my project </h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact page</h1><p>To get in touch, email me at <a href=\"mailto:alshahadath@gmail.com\">alshahadath@gmail.com</a></p>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		pageNotFoundHandler(w, r)
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on :3000...")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		panic(err)
	}
}
