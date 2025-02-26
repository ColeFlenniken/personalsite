package main

import (
	"fmt"
	"net/http"

	"github.com/ColeFlenniken/personalsite/API"
	"github.com/ColeFlenniken/personalsite/Templates"
	"github.com/a-h/templ"
)

func main() {
	component := Templates.MainLayout() // Write the response

	http.Handle("/", templ.Handler(component))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("./static/css"))))
	http.Handle("/static/scripts/", http.StripPrefix("/static/scripts/", http.FileServer(http.Dir("./static/scripts"))))
	http.HandleFunc("/Card", func(w http.ResponseWriter, r *http.Request) {
		API.GetJobCard(w, r)
	})

	http.HandleFunc("/Test", func(w http.ResponseWriter, r *http.Request) {
		API.GetTest(w, r)
	})
	http.HandleFunc("/Resume", func(w http.ResponseWriter, r *http.Request) {
		API.GetResume(w, r)
	})

	http.HandleFunc("/Index", func(w http.ResponseWriter, r *http.Request) {
		API.GetIndex(w, r)
	})

	http.HandleFunc("/Index/Letters", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HERE")
		API.GetNextLetter(w, r)
	})
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
