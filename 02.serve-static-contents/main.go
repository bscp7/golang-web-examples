package main

import (
	"fmt"
	"net/http"
)

func rootRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `Static content server
## Help ##
Copy the binary and place it with a directory named "static".
Add static contents in the "static" directory and navigate to http://localhost:8080/static/.
`)
}

func main() {
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	//Handlers
	http.HandleFunc("/", rootRequestHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//Start server
	panic(http.ListenAndServe(addr, nil))
}
