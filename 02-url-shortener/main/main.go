package main

import (
	"fmt"
	"net/http"

	// To install a local package we take the name of the current
	// module (github.com/emilioschepis/gophercises) and append the
	// name of the directory which contains the package
	// In this case my `handler.go` file is in 02-url-shortener
	// In a real project I think that the name of the folder would
	// need (or at least it would be good practice) to match the
	// name of the package
	urlshort "github.com/emilioschepis/gophercises/02-url-shortener"
)

func main() {
	// A `mux` is short for "multiplexer" and its role is to dispatch
	// requests to the appropriate handlers
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
