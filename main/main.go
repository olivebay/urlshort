package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/olivebay/urlshort"
)

func main() {
	//Pass the file path as a command line flag
	yamlFlag := flag.String("yaml", "", "Specify yaml file")
	flag.Parse()

	mux := defaultMux()

	//Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlFile, err := ioutil.ReadFile(*yamlFlag)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	yamlHandler, err := urlshort.YAMLHandler([]byte(yamlFile), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8888")
	http.ListenAndServe(":8888", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
