package testutil

import "io/ioutil"

// Fixture loads a fixture file to a string
func Fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
