package main

import (
	"fmt"
	"net/http"

	"github.com/ColeFlenniken/personalsite/templates"
	"github.com/a-h/templ"
)

func main() {
	component := templates.MainPage()

	http.Handle("/", templ.Handler(component))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("./static/css"))))
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
