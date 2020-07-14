// This file is separated from the others as it is a good practice to put all commands in a `cmd/<name>` directory
package main

import (
	"flag"
	"fmt"
	"os"

	cyoa "github.com/emilioschepis/gophercises/03-choose-your-own-adventure"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		// Usually `panic`-ing is not a great idea but it works for quick tries
		panic(err)
	}

	story, err := cyoa.JsonStory(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
