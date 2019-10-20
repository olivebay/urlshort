package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/olivebay/urlshort"
)

func main() {
	// Pass the file path as a command line flag
	yamlFlag := flag.String("yaml", "", "Specify a YAML file")
	jsonFlag := flag.String("json", "", "Specify a JSON file")
	flag.Parse()

	var fileHandler http.HandlerFunc

	if *yamlFlag != "" {
		yamlFile, err := ioutil.ReadFile(*yamlFlag)
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}
		fileHandler, err = urlshort.YAMLHandler([]byte(yamlFile))
		if err != nil {
			panic(err)
		}
	}

	if *jsonFlag != "" {
		jsonFile, err := ioutil.ReadFile(*jsonFlag)
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}
		fileHandler, err = urlshort.JSONHandler([]byte(jsonFile))
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Starting the server on :8888")
	http.ListenAndServe(":8888", fileHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
