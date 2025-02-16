package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := mainpage()

	http.Handle("/", templ.Handler(component))
	http.Handle("/css/", http.FileServer(http.Dir("./")))
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
