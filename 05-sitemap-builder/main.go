package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	link "github.com/emilioschepis/gophercises/04-html-link-parser"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	// Named urlFlag in order to avoid conflict with the url package
	urlFlag := flag.String("url", "https://emilioschepis.com/", "the url that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 3, "the maximum number of links to traverse")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)
	found := urlset{
		Xmlns: xmlns,
	}
	for _, page := range pages {
		found.Urls = append(found.Urls, loc{page})
	}

	// Print the constant header of the XML package before the rest of the urlset
	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(found); err != nil {
		panic(err)
	}
	fmt.Println()
}

func bfs(urlStr string, maxDepth int) []string {
	// Having the value of the map as `struct{}` represents a simple way to reap all the benefits
	// of using a map (access in O(1)) while not allocating much space (an empty struct is the smallest type in Go).
	seen := make(map[string]struct{})

	var q map[string]struct{}
	nq := map[string]struct{}{
		// This used to be `struct{}{}` but the "simplifycompositelit" rule suggests to turn this into a simple literal: `{}`.
		urlStr: {},
	}

	for i := 0; i < maxDepth; i++ {
		// When we are done with a given queue, we copy nq (next queue) into it and reinitialize nq.
		q, nq = nq, make(map[string]struct{})

		// Exit early if there are no more elements queued.
		if len(q) == 0 {
			break
		}

		// for key, value := ...
		for url := range q {
			// Ok tells us if `seen` has this key or not.
			if _, ok := seen[url]; ok {
				continue
			}

			// Otherwise add the url to the `seen` map (with a value of empty struct).
			seen[url] = struct{}{}

			for _, link := range get(url) {
				if _, ok := seen[link]; !ok {
					// Add each found link to the `nq` if they are new.
					nq[link] = struct{}{}
				}
			}
		}
	}

	// Preallocate a slice with a capacity of the amount of found hrefs.
	ret := make([]string, 0, len(seen))
	for url := range seen {
		ret = append(ret, url)
	}
	return ret
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
