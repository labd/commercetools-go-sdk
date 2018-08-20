package testutil

import "io/ioutil"

func Fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
