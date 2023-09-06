package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/talksik/online-status-htmx-golang/templates"
)

var (
	statusOne string
	statusTwo string
)

func getStatusHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.GetStatus(statusOne, statusTwo)
	component.Render(r.Context(), w)
}

func updateStatusHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: read request and update statuses

	println("changing status")

  queryParams := r.URL.Query()
	h_status := queryParams.Get("statush")
	a_status := queryParams.Get("statusa")

  statusOne = h_status
  statusTwo = a_status

	getStatusHandler(w, r)
}

func main() {
	println("yo")

	component := templates.Hello("John")
	http.Handle("/hello", templ.Handler(component))

	http.HandleFunc("/changestatus", updateStatusHandler)
	http.HandleFunc("/", getStatusHandler)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
