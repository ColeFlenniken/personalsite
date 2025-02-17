package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ColeFlenniken/personalsite/templates"
	"github.com/a-h/templ"
)

/*


func main() {
    http.HandleFunc("/example", func(res http.ResponseWriter, req *http.Request) {
        queryParamDisplayHandler(res, req)
    })
    println("Enter this in your browser:  http://localhost:8080/example?name=jenny&phone=867-5309")
    http.ListenAndServe(":8080", nil)
}

*/

func queryParamDisplayHandler(title string, description string) templ.Component {

	return templates.FrontPageCard(title, description)

}

func main() {
	component := templates.MainPage() // Write the response

	http.Handle("/", templ.Handler(component))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("./static/css"))))
	http.HandleFunc("/Card", func(w http.ResponseWriter, r *http.Request) {
		// Parse the query parameters
		query := r.URL.Query()
		//need to add error handling
		id, err := strconv.Atoi(query.Get("card"))
		if err != nil {
			fmt.Println(err)
		}
		titles := [3]string{"Carda", "cardb", "cardc"}
		descs := [3]string{"desca", "descb", "descc"}
		queryParamDisplayHandler(titles[id], descs[id]).Render(r.Context(), w)
	})
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
