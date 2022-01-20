package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request")
	url := "https://www.random.org/integers/?num=3&min=0&max=3&col=1&base=10&format=plain"
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err)

	}
	fmt.Print(responseData)
}

func main() {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/", getRandom)
	log.Fatal(http.ListenAndServe(":8080", router))
}
