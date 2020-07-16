package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	link "github.com/emilioschepis/gophercises/04-html-link-parser"
)

func main() {
	// Named urlFlag in order to avoid conflict with the url package
	urlFlag := flag.String("url", "https://emilioschepis.com/", "the url that you want to build a sitemap for")
	flag.Parse()

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	links, err := link.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	reqURL := resp.Request.URL

	// Create a fresh URL using only the scheme and the host (ignore trailing /, paths, etc.)
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}

	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, baseURL.String()+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			// This case obviously matches https too
			hrefs = append(hrefs, l.Href)
		default:
			continue
		}
	}

	for _, href := range hrefs {
		fmt.Println(href)
	}
}
