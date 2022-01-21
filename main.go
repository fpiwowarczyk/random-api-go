package main

import (
	"net/http"

	"github.com/fpiwowarczyk/nobl9-go/random"
)

func setUpApi() {

	http.HandleFunc("/random/mean", random.GetRandomValues)

	http.ListenAndServe(":8080", nil)
}

func main() {
	setUpApi()

}
