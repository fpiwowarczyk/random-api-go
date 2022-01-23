package random

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatResponses_Success(t *testing.T) {
	testData := [][]int{{1, 2}}
	expected := []RandomValuesResponse{{Stddev: 1, Data: []int{1, 2}}, {Stddev: 1, Data: []int{1, 2}}}

	output, err := FormatResponses(testData)

	assert.Nil(t, err, "Shoudn't return any errors")
	assert.NotNil(t, output, "Valid data should return data")
	assert.Equal(t, expected, output, "Returned data should contain input with standard deviation and result for merged values")
}

func TestFormatResponses_Fail(t *testing.T) {
	testData := [][]int{{}}
	output, err := FormatResponses(testData)

	assert.Nil(t, output)
	assert.Equal(t, err, errors.New("No data present for formatting"))
}

func TestCountStandardDeviation_Success(t *testing.T) {
	testData := []int{1, 2, 3, 4}

	output, err := CountStandardDeviation(testData)

	assert.Nil(t, err, "No error should be returned")
	assert.NotNil(t, output, "Some data should be returned")
	assert.Equal(t, output, 1.4142135623730951, "Should return correct data")
}

func TestCountStandardDeviation_Fail(t *testing.T) {
	testData := []int{}

	_, err := CountStandardDeviation(testData)

	assert.Equal(t, err, errors.New("Missing data for counting standard deviation"), "Should return correct data")
}

func TestConvertByteesToIntegers_Success(t *testing.T) {
	testData := []byte{49, 10, 50, 10}

	output, err := ConvertBytesToIntegers(testData)

	assert.Nil(t, err, "No errors for valid data")
	assert.Equal(t, []int{1, 2}, output, "Converted data should be equal to int slice 1,2")
}

func TestConvertByteesToIntegers_Fail(t *testing.T) {
	testData := []byte{}

	output, err := ConvertBytesToIntegers(testData)

	assert.Equal(t, errors.New("Missing data for converting"), err, "No errors for valid data")
	assert.Nil(t, nil, output, "Output should be nil")
}
