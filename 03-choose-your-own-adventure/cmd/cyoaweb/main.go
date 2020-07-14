// This file is separated from the others as it is a good practice to put all commands in a `cmd/<name>` directory
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	cyoa "github.com/emilioschepis/gophercises/03-choose-your-own-adventure"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		// Usually `panic`-ing is not a great idea but it works for quick tries
		panic(err)
	}

	story, err := cyoa.JSONStory(file)
	if err != nil {
		panic(err)
	}

	// We can pass dynamic options!
	h := cyoa.NewHandler(story /*, cyoa.WithTemplate(nil)*/)
	fmt.Printf("Starting the server on port: %d\n", *port)

	// The HTTP server should always be kept alive
	// If it returns for some reason, the reason will be automatically logged
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
