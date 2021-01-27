package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

//ResponsePayload struct defines response object
//Parameters in accordance with templates/response.html
type ResponsePayload struct {
	Host       string
	RequestURI string
	Time       string
}

func rootRequestHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/response.html"))

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	data := &ResponsePayload{
		Host:       hostname,
		RequestURI: r.URL.RequestURI(),
		Time:       time.Now().Format(time.RFC3339),
	}
	tmpl.Execute(w, data)
}

func main() {
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	//Handlers
	http.HandleFunc("/", rootRequestHandler)
	//Start server
	panic(http.ListenAndServe(addr, nil))
}
