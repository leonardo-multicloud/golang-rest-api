package main

import (
	"log"
	"net/http"

	"github.com/leonardo-multicloud/golang-rest-api/api"
)

func main() {

	economy := api.Economy{
		HTTPClient: http.DefaultClient,
	}

	http.HandleFunc("/price", economy.GetPrice)
	log.Println("Server Listen...")
	panic(http.ListenAndServe("0.0.0.0:8000", nil))
}
