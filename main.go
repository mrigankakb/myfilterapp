package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	jsonfilter "github.com/digitapai/jsonfilter/filter"
	"github.com/yalp/jsonpath"
)

func main() {
	var (
		dataFile   string
		filterFile string
	)

	flag.StringVar(&dataFile, "d", "", "")
	flag.StringVar(&filterFile, "f", "", "")
	flag.Parse()

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
}

func writeOutput(v interface{}) {
	if b, err := json.Marshal(v); err == nil {
		var out bytes.Buffer
		if err = json.Indent(&out, b, "", "  "); err == nil {
			writer := bufio.NewWriter(os.Stdout)
			if _, err = out.WriteTo(writer); err == nil {
				err = writer.Flush()
			}
		}
	}

}

//ApplyFilter ... apply a filter on the input json
//the filter json should be a matching subtree of the input JSON
func ApplyFilter(input []byte, filter []byte) (interface{}, error) {
	v, err := jsonfilter.FilterJsonFromTextWithFilterRunner(string(input), string(filter), func(command string, value string) (string, error) {

		return value, nil
	})

	return v, err
}

//RunTemplate ... create a output as per template.
//The template will contain a jsonpath based expression for each field.
//This json path expression will be applied to the input JSON to read the value for the field.
func RunTemplate(template []byte, input []byte) (interface{}, error) {

	var data interface{}
	json.Unmarshal(input, &data)

	v, err := jsonfilter.FilterJsonFromTextWithFilterRunner(string(template), string(template), func(command string, value string) (string, error) {
		filterval, _ := jsonpath.Read(data, value)
		// fmt.Println(filterval)
		ret := fmt.Sprintf("%v", filterval)
		return ret, nil
	})

	return v, err

}
