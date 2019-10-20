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

	mux := defaultMux()

	var fileHandler http.HandlerFunc

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	if *yamlFlag != "" {
		yamlFile, err := ioutil.ReadFile(*yamlFlag)
		if err != nil {
			fmt.Println("File reading error", err)
			return
		}
		fileHandler, err = urlshort.YAMLHandler([]byte(yamlFile), mapHandler)
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
		fileHandler, err = urlshort.JSONHandler([]byte(jsonFile), mapHandler)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Starting the server on :8888")
	http.ListenAndServe(":8888", fileHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
