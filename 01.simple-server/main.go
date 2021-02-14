package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//ResponsePayload struct defines response object
//Parameters in accordance with templates/response.html
type ResponsePayload struct {
	Host       string
	RequestURI string
	Time       string
	Version    string
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
		Version:    "v2",
	}
	tmpl.Execute(w, data)
}

func main() {
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	http.HandleFunc("/", rootRequestHandler)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			panic(err)
		}
	}()
	log.Println("Server started...")

	<-done
	log.Print("Server Stopped")
}
