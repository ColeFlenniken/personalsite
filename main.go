package main

import (
	"fmt"
	"net/http"

	"github.com/ColeFlenniken/personalsite/API"
	"github.com/ColeFlenniken/personalsite/Templates"
	"github.com/a-h/templ"
)

func main() {

	http.HandleFunc("/saveCanvas", func(w http.ResponseWriter, r *http.Request) {
		API.UpdateCanvasData(w, r)

	})
	http.HandleFunc("/getImageData", func(w http.ResponseWriter, r *http.Request) {
		API.GetCanvasPixels(w, r)
	})

	http.Handle("/", templ.Handler(Templates.MainLayout()))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("./static/css"))))
	http.Handle("/static/scripts/", http.StripPrefix("/static/scripts/", http.FileServer(http.Dir("./static/scripts"))))
	http.Handle("/static/img/", http.StripPrefix("/static/img/", http.FileServer(http.Dir("./static/img"))))

	http.HandleFunc("/Resume", func(w http.ResponseWriter, r *http.Request) {
		API.GetResume(w, r)
	})

	http.HandleFunc("/Index", func(w http.ResponseWriter, r *http.Request) {
		API.GetIndex(w, r)
	})
	http.HandleFunc("/Resume/Education", func(w http.ResponseWriter, r *http.Request) {
		API.GetEducation(w, r)
	})
	http.HandleFunc("/Resume/Jobs", func(w http.ResponseWriter, r *http.Request) {
		API.GetJobs(w, r)
	})
	http.HandleFunc("/About", func(w http.ResponseWriter, r *http.Request) {
		API.GetAbout(w, r)
	})
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
