package main

import (
	"fmt"
	"net/http"
)
type PageFunc func(http.ResponseWriter, *http.Request)
type WebPage struct {
	Title string
	Home PageFunc
	About PageFunc
	Contact PageFunc
}
var page WebPage
func main() {
	mux:= http.NewServeMux()
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		if page.Home == nil {
			fmt.Fprintf(w, "Home")
		} else {
			page.Home(w, r)
		}
	})
	mux.HandleFunc("/about", func (w http.ResponseWriter, r *http.Request) {
		if page.About == nil {
			fmt.Fprintf(w, "About")
		} else {
			page.About(w, r)
		}
	})
	mux.HandleFunc("/contact", func (w http.ResponseWriter, r *http.Request) {
		if page.Contact == nil {
			fmt.Fprintf(w, "Contact")
		} else {
			page.Contact(w, r)
		}
	})
	http.ListenAndServe(":8080", mux)
}

