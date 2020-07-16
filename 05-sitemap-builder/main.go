package main

import (
	"flag"
)

func main() {
	// Named urlFlag in order to avoid conflict with the url package
	urlFlag := flag.String("url", "https://emilioschepis.com/", "the url that you want to build a sitemap for")
	flag.Parse()

	_ = urlFlag
}
