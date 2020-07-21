package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	apiFile := os.Args[1]
	fmt.Printf("Generating code for file %s\n", apiFile)
	fi, err := os.Stat(apiFile)
	if err != nil {
		panic(err)
	}
	size := fi.Size()
	if size == 0 {
		panic(fmt.Errorf("File %s is empty", apiFile))
	}

	f, err := ioutil.ReadFile(apiFile)
	if err != nil {
		panic(err)
	}

	t := yaml.MapSlice{}

	err = yaml.Unmarshal(f, &t)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Println("=================[ Parsing RAML ]=================")
	objects, resources := parseYaml(t)
	postProcess(objects)

	log.Println("===============[ Generating types ]===============")
	generateTypes(objects)
	log.Println("==============[ Generating services ]=============")
	generateServices(objects, resources)
}
