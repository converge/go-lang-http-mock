package main

import (
	"gohttpmock/api"
	"log"
	"net/http"
)

func main() {


	service := api.Service {
		HTTPClient: http.DefaultClient,
	}

	http.HandleFunc("/", service.GetRandomUser)
	log.Println("listening...")
	panic(http.ListenAndServe("localhost:7000", nil))
}