package parser

import (
	"testing"

	"fmt"
	"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("test_parser.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)
	const resultSize = 470

	if len(result.Requests) != resultSize {
		fmt.Printf("request number errr, should be %d, but got %d", resultSize, result.Requests)
	}
	if len(result.Items) != resultSize {
		fmt.Printf("request number errr, should be %d, but got %d", resultSize, result.Items)
	}
}
