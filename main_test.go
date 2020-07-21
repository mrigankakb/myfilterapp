package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/yalp/jsonpath"
)

//TestFilter ... test fixed template
func TestFilter(t *testing.T) {
	dataFile := "test/data/input.json"
	filterFile := "test/data/filter.json"

	df, err := os.Open(dataFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer df.Close()

	data, err := ioutil.ReadAll(df)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ff, err := os.Open(filterFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer ff.Close()

	filter, err := ioutil.ReadAll(ff)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	v, err := ApplyFilter(data, filter)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	writeOutput(v)
	fmt.Println("")
}

//TestTemplate ... test dynamic template
func TestTemplate(t *testing.T) {
	templateFile := "test/data/template.json"
	inputFile := "test/data/input.json"

	tf, err := os.Open(templateFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer tf.Close()

	template, err := ioutil.ReadAll(tf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	inf, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer inf.Close()

	input, err := ioutil.ReadAll(inf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	v, err := RunTemplate(template, input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	writeOutput(v)
	fmt.Println("")
}

func TestJSONPath(t *testing.T) {
	inputFile := "test/data/input.json"

	inf, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer inf.Close()

	input, err := ioutil.ReadAll(inf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var data interface{}
	json.Unmarshal(input, &data)

	v, _ := jsonpath.Read(data, "$.ENTITY.CASA[0]")

	fmt.Println(v)
}
