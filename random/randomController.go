package random

import (
	"encoding/json"
	"log"
	"net/http"
)

type RandomValuesResponse struct {
	Stddev float64 `json:"stddev"`
	Data   []int   `json:"data"`
}

// Get random numbers and handle wrong input values
func GetRandomValues(w http.ResponseWriter, r *http.Request) {
	log.Printf("Call for endpoint: %s", r.URL.String())

	requests, length, err := ParseQueryParams(w, r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	numbers, err := GetNumbers(requests, length)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusRequestTimeout)
		return
	}

	response := FormatResponses(numbers)

	json.NewEncoder(w).Encode(response)
}
