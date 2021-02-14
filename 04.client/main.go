package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Payload struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	data := &Payload{
		Title:  "foo",
		Body:   "bar",
		UserId: 1,
	}
	jData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	payload := bytes.NewReader(jData)
	request, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(resp)
	}

	fmt.Println(resp)
}
