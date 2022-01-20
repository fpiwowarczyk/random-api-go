package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type randomResponse struct {
	Stddev float64 `json:"stddev"`
	Data   []int64 `json:"data"`
}

var tempResponse = []randomResponse{
	{Stddev: 1, Data: []int64{1, 2, 3, 4}},
	{Stddev: 1, Data: []int64{2, 3, 4, 5}},
	{Stddev: 1, Data: []int64{1, 2, 2, 3, 3, 4, 4, 5}},
}

var tempReturnedData = []int{1, 2, 3, 4, 5}

// get random respons with the deviation and values
func getRandom(c *gin.Context) {
	// Get data
	var mean = 0
	for i := 0; i < len(tempReturnedData); i++ {
		mean += tempReturnedData[i]
	}
	mean = mean / len(tempResponse)
	fmt.Print(mean)
	c.IndentedJSON(http.StatusOK, tempResponse)
}

func main() {
	router := gin.Default()
	router.GET("/random/", getRandom)

	router.Run("localhost:8080")
}
