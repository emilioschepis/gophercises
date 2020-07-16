package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	link "github.com/emilioschepis/gophercises/04-html-link-parser"
)

func main() {
	// Named urlFlag in order to avoid conflict with the url package
	urlFlag := flag.String("url", "https://emilioschepis.com/", "the url that you want to build a sitemap for")
	flag.Parse()

	pages := get(*urlFlag)
	for _, href := range pages {
		fmt.Println(href)
	}
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqURL := resp.Request.URL

	// Create a fresh URL using only the scheme and the host (ignore trailing /, paths, etc.)
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()

	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(r io.Reader, base string) []string {
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}

	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			// This case obviously matches https too
			ret = append(ret, l.Href)
		default:
			continue
		}
	}

	return ret
}

func filter(links []string, predicate func(string) bool) []string {
	var ret []string

	for _, link := range links {
		if predicate(link) {
			ret = append(ret, link)
		}
	}

	return ret
}

func withPrefix(p string) func(string) bool {
	return func(s string) bool {
		return strings.HasPrefix(s, p)
	}
}
