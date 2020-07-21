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
)

var (
	dataFile   string
	filterFile string
)

func main() {

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

	data = data
	filter = filter

	v, err := jsonfilter.FilterJsonFromText(string(data), string(filter))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

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

//CreateOutputJSON ... create a output as per client template
func CreateOutputJSON(template []byte, input []byte) []byte {

	return nil
}
