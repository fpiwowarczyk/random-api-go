package random

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type RandomResponse struct {
	Stddev float64 `json:"stddev"`
	Data   []int64 `json:"data"`
}

type Result struct {
	FirstNumbers  *[]int `json:"firstNumbers,omitempty"`
	SecondNumbers *[]int `json:"secondNumbers,omitempty"`
}

// Get random numbers and handle wrong input values
func GetRandomValues(w http.ResponseWriter, r *http.Request) {
	requestsStr, ok := r.URL.Query()["requests"]
	requests, err := strconv.Atoi(requestsStr[0])
	if !ok || err != nil || len(requestsStr) > 1 {
		log.Printf("Bad request for value requests: %v", requestsStr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lengthStr, ok := r.URL.Query()["length"]
	length, err := strconv.Atoi(lengthStr[0])
	if !ok || err != nil || len(lengthStr) > 1 {
		log.Printf("Bad request for value length: %v", lengthStr)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Requests: %d Length: %d", requests, length)

	response := getNumbers(requests, length)

	json.NewEncoder(w).Encode(response)
}

func getNumbers(req int, len int) [][]int {
	var responses [][]int

	return responses

}
