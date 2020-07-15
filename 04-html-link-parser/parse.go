package link

import (
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

	nodes := linkNodes(doc)
	var links []Link

	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}

	ret.Text = "TODO"
	return ret
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		// If the passed element is an anchor, return it directly.
		return []*html.Node{n}
	}

	var ret []*html.Node
	// This is just a normal for loop, the third part simply instructs the loop on
	// what to do after each iteration (select the next sibling node).
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// The linkNodes function returns a slice, so the ... (variadic) allows us to append each element separately.
		// For example one slice with ten elements becomes ten different arguments for the function
		ret = append(ret, linkNodes(c)...)
	}

	return ret
}
