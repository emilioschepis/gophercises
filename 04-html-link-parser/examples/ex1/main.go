package main

import (
	"fmt"
	"strings"

	link "github.com/emilioschepis/gophercises/04-html-link-parser"
)

var exampleHTML = `
<html>
	<body>
		<h1>Hello!</h1>
		<a href="/other-page">A link to another page</a>
		<a href="/page-two">A link to a second page <span>with a span</span></a>
	</body>
</html>
`

func main() {
	// Creates a `io.Reader` from a string.
	r := strings.NewReader(exampleHTML)

	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
