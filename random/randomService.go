package random

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

func ParseQueryParams(w http.ResponseWriter, r *http.Request) (int, int, error) {
	requestsStr, ok := r.URL.Query()["requests"]
	requests, err := strconv.Atoi(requestsStr[0])
	if !ok || err != nil || len(requestsStr) > 1 {
		return 0, 0, errors.New("Bad requests value")
	}
	lengthStr, ok := r.URL.Query()["length"]
	length, err := strconv.Atoi(lengthStr[0])
	if !ok || err != nil || len(lengthStr) > 1 {
		return 0, 0, errors.New("Bad length value")
	}
	return requests, length, nil

}

// Dont call it with more than 5 requests or youll get banned :(
func GetNumbers(req int, len int) ([][]int, error) {
	var wg sync.WaitGroup
	var chanArray []chan []int
	var chanError []chan error
	var allValues [][]int

	for i := 0; i < req; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		chanArray = append(chanArray, make(chan []int))
		chanError = append(chanError, make(chan error))

		go func(id int, ctx context.Context, ch chan []int, errCh chan error) {
			wg.Add(1)
			defer close(errCh)
			defer close(ch)
			defer wg.Done()

			req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://www.random.org/integers/?num=%d&min=1&max=100&col=1&base=10&format=plain&rnd=new", len), nil)
			if err != nil {
				errCh <- err
				log.Println(err)
				return
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				errCh <- err
				log.Println(err)
				return
			}

			defer resp.Body.Close()
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				errCh <- err
				log.Println(err)
				return
			}

			values, err := ConvertBytesToIntegers(bodyBytes)
			if err != nil {
				errCh <- err
				log.Println(err)
				return
			}
			ch <- values
			return
		}(i, ctx, chanArray[i], chanError[i])
	}

	wg.Wait()

	for i := 0; i < req; i++ {
		select {
		case values := <-chanArray[i]:
			allValues = append(allValues, values)
		case err := <-chanError[i]:
			return nil, err
		}
	}
	return allValues, nil
}

func FormatResponses(numbers [][]int) []RandomValuesResponse {
	var response []RandomValuesResponse
	var sumOfAllValues []int
	for _, data := range numbers {
		sumOfAllValues = append(sumOfAllValues, data...)
		response = append(response, RandomValuesResponse{CountStandardDeviation(data), data})
	}
	response = append(response, RandomValuesResponse{CountStandardDeviation(sumOfAllValues), sumOfAllValues})

	return response
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
	return math.Sqrt(float64(toMultiply) * squaredSums)
}

func ConvertBytesToIntegers(bodyBytes []byte) ([]int, error) {
	var values []int
	valuesStr := strings.Split(string(bodyBytes), "\n")
	for _, i := range valuesStr {
		if i != "" {
			value, err := strconv.Atoi(i)
			if err != nil {
				return nil, err
			}
			values = append(values, value)
		}
	}

	return values, nil
}
