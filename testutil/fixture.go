package testutil

import "io/ioutil"

// Fixture loads a fixture file to a string
func Fixture(path string, statusCode int) ResponseData {
	b, err := ioutil.ReadFile("fixtures/" + path)
	if err != nil {
		panic(err)
	}
	return ResponseData{StatusCode: statusCode, Body: string(b)}
}
