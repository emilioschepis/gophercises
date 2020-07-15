package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents a link in an HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and return a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	dfs(doc, "")

	return nil, nil
}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}

	fmt.Println(padding, msg)

	// This is just a normal for loop, the third part simply instructs the loop on
	// what to do after each iteration (select the next sibling node).
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
