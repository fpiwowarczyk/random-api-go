package random

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type RandomValuesResponse struct {
	Stddev float64 `json:"stddev"`
	Data   []int   `json:"data"`
}

// Get random numbers and handle wrong input values
func GetRandomValues(w http.ResponseWriter, r *http.Request) {
	log.Printf("Call for endpoint: %s", r.URL.String())
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

	randomValues := getNumbers(requests, length)

	var responses []RandomValuesResponse
	var sumOfAllValues []int
	for _, data := range randomValues {
		sumOfAllValues = append(sumOfAllValues, data...)
		responses = append(responses, RandomValuesResponse{CountStandardDeviation(data), data})
	}
	responses = append(responses, RandomValuesResponse{CountStandardDeviation(sumOfAllValues), sumOfAllValues})
	json.NewEncoder(w).Encode(responses)
}

func CountStandardDeviation(values []int) float64 {
	mean := 0
	for _, i := range values {
		mean = mean + i
	}
	mean = mean / len(values)
	toMultiply := 1.0 / (float64(len(values)) - 1.0)
	squaredSums := 0.0
	for _, i := range values {
		squaredSums = squaredSums + math.Pow(float64(i-mean), 2)
	}
	log.Println(float64(toMultiply))
	return math.Sqrt(float64(toMultiply) * squaredSums)
}

// Dont call it with more than 5 requests or youll get banned :(
func getNumbers(req int, len int) [][]int {
	var wg sync.WaitGroup
	var chanArray []chan []int
	var allValues [][]int

	for i := 0; i < req; i++ {
		chanArray = append(chanArray, make(chan []int))
		go func(id int, ch chan []int) {
			wg.Add(1)
			log.Printf("Calling goroutine with id:%d \n", id)
			resp, err := http.Get(fmt.Sprintf("https://www.random.org/integers/?num=%d&min=1&max=100&col=1&base=10&format=plain&rnd=new", len))
			if err != nil {
				log.Println(err)
			}
			if resp.StatusCode == http.StatusOK {
				defer resp.Body.Close()
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Println(err)
				}
				var values []int

				valuesStr := strings.Split(string(bodyBytes), "\n")
				for _, i := range valuesStr {
					if i != "" {
						value, err := strconv.Atoi(i)
						if err != nil {
							log.Println(err)
						}
						values = append(values, value)
					}
				}
				ch <- values
			}
			wg.Done()
		}(i, chanArray[i])
	}

	wg.Wait()
	for i := 0; i < req; i++ {
		allValues = append(allValues, <-chanArray[i])
	}
	return allValues
}
