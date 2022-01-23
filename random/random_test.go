package random

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestFormatResponses_Success(t *testing.T) {
// 	//given
// 	testData := [][]int{{1, 2}}
// 	expected := []RandomValuesResponse{{Stddev: 1, Data: []int{1, 2}}, {Stddev: 1, Data: []int{1, 2}}}

// 	//when
// 	output, err := FormatResponses(testData)

// 	//then
// 	assert.Nil(t, err, "Shoudn't return any errors")
// 	assert.NotNil(t, output, "Valid data should return data")
// 	assert.Equal(t, expected, output, "Returned data should contain input with standard deviation and result for merged values")
// }

func TestFormatResponses_Fail(t *testing.T) {
	//given
	testData := [][]int{{}}

	//when
	output, err := FormatResponses(testData)

	//then
	assert.Nil(t, output)
	assert.Equal(t, err, errors.New("No data present for formatting"))
}

// func TestCountStandardDeviation_Success(t *testing.T) {
// 	//given
// 	testData := []int{1, 2, 3, 4}

// 	//when
// 	output, err := CountStandardDeviation(testData)

// 	//then
// 	assert.Nil(t, err, "No error should be returned")
// 	assert.NotNil(t, output, "Some data should be returned")
// 	assert.Equal(t, output, 1.4142135623730951, "Should return correct data")
// }

// func TestCountStandardDeviation_Fail(t *testing.T) {
// 	//given
// 	testData := []int{}

// 	//when
// 	_, err := CountStandardDeviation(testData)

// 	//then
// 	assert.Equal(t, err, errors.New("Missing data for counting standard deviation"), "Should return correct data")
// }

// func TestConvertByteesToIntegers_Success(t *testing.T) {
// 	//given
// 	testData := []byte{49, 10, 50, 10}

// 	//when
// 	output, err := ConvertBytesToIntegers(testData)

// 	//then
// 	assert.Nil(t, err, "No errors for valid data")
// 	assert.Equal(t, []int{1, 2}, output, "Converted data should be equal to int slice 1,2")
// }

// func TestConvertByteesToIntegers_Fail(t *testing.T) {
// 	//given
// 	testData := []byte{}

// 	//when
// 	output, err := ConvertBytesToIntegers(testData)

// 	//then
// 	assert.Equal(t, errors.New("Missing data for converting"), err, "No errors for valid data")
// 	assert.Nil(t, nil, output, "Output should be nil")
// }

// func TestParseQueryParams_success(t *testing.T) {
// 	handler := func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "Some basic output")
// 	}

// 	req := httptest.NewRequest("GET", "https://random.org/foo?requests=1&length=2", nil)
// 	w := httptest.NewRecorder()
// 	handler(w, req)

// 	requests, length, err := ParseQueryParams(w, req)

// 	assert.Nil(t, err, "No errors for valid values")
// 	assert.Equal(t, 1, requests, "Should return valid number of requests")
// 	assert.Equal(t, 2, length, "Should return valid length")
// }

// func TestParseQueryParams_Fail_ToManyValues(t *testing.T) {
// 	handler := func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "Some basic output")
// 	}

// 	req := httptest.NewRequest("GET", "https://random.org/foo?requests=1,2&length=2", nil)
// 	w := httptest.NewRecorder()
// 	handler(w, req)

// 	_, _, err := ParseQueryParams(w, req)

// 	assert.NotNil(t, err, "Error should be present for multiple values")
// 	assert.Equal(t, errors.New("Bad requests value"), err, "Should return given error")
// }

// func TestParseQueryParams_Fail_NoANumberValue(t *testing.T) {
// 	handler := func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "Some basic output")
// 	}

// 	req := httptest.NewRequest("GET", "https://random.org/foo?requests=1&length=a", nil)
// 	w := httptest.NewRecorder()
// 	handler(w, req)

// 	_, _, err := ParseQueryParams(w, req)

// 	assert.NotNil(t, err, "Error should be present for multiple values")
// 	assert.Equal(t, errors.New("Bad length value"), err, "Should return given error")
// }

// func TestGetNumbers_Success(t *testing.T) {
// 	len := 2
// 	req := 1
// 	httpmock.Activate()
// 	defer httpmock.DeactivateAndReset()

// 	httpmock.RegisterResponder("GET", fmt.Sprintf("https://www.random.org/integers/?num=%d&min=1&max=100&col=1&base=10&format=plain&rnd=new", len),
// 		func(req *http.Request) (*http.Response, error) {
// 			resp := httpmock.NewBytesResponse(200, []byte{49, 10, 50, 10}) // Bytes response of [1,2]
// 			return resp, nil
// 		})

// 	result, err := GetNumbers(req, len)

// 	assert.Nil(t, err, "No errors for valid data")
// 	assert.Equal(t, [][]int{{1, 2}}, result, "Result data should be queal to expected")
// }

// func TestGetNumbers_Success_Multiple_Requests(t *testing.T) {
// 	len := 2
// 	req := 2
// 	httpmock.Activate()
// 	defer httpmock.DeactivateAndReset()

// 	httpmock.RegisterResponder("GET", fmt.Sprintf("https://www.random.org/integers/?num=%d&min=1&max=100&col=1&base=10&format=plain&rnd=new", len),
// 		func(req *http.Request) (*http.Response, error) {
// 			resp := httpmock.NewBytesResponse(200, []byte{49, 10, 50, 10}) // Bytes response of [1,2]
// 			return resp, nil
// 		})

// 	result, err := GetNumbers(req, len)

// 	assert.Nil(t, err, "No errors for valid data")
// 	assert.Equal(t, [][]int{{1, 2}, {1, 2}}, result, "Result data should be queal to expected")
// }

// func TestGetNumbers_Fail_Timeout(t *testing.T) {
// 	len := 2
// 	req := 1
// 	expectedError := url.Error{Op: "Get", URL: "https://www.random.org/integers/?num=2&min=1&max=100&col=1&base=10&format=plain&rnd=new", Err: context.DeadlineExceeded}
// 	httpmock.Activate()
// 	defer httpmock.DeactivateAndReset()

// 	httpmock.RegisterResponder("GET", fmt.Sprintf("https://www.random.org/integers/?num=%d&min=1&max=100&col=1&base=10&format=plain&rnd=new", len),
// 		func(req *http.Request) (*http.Response, error) {
// 			time.Sleep(time.Second * 3)
// 			resp := httpmock.NewBytesResponse(200, []byte{49, 10, 50, 10}) // Bytes response of [1,2]
// 			return resp, nil
// 		})

// 	result, err := GetNumbers(req, len)

// 	fmt.Println(err)
// 	assert.Equal(t, &expectedError, err, "Error for timeout")
// 	assert.Nil(t, result, "Data should be nil for timeout")
// }
