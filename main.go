package main

import (
	"log"
	"net/http"

	"github.com/fpiwowarczyk/random-api-go/random"
)

var PORT = "8080"

func setUpApi() {

	http.HandleFunc("/random/mean", random.GetRandomValues)

	log.Printf("Server is listening at port %s", PORT)
	log.Fatalln(http.ListenAndServe(":"+PORT, nil))
}

func main() {
	setUpApi()
}
