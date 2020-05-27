package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pratikms/gophercises/linkparser"
)

func hrefs(body io.Reader, base string) []string {
	var hrefs []string
	links, _ := linkparser.Parse(body)
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}

	return hrefs
}

func get(urlStr string) []string {
	fmt.Println(urlStr)
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()
	fmt.Println("Request URL: ", reqURL)
	fmt.Println("Base URL: ", base)

	return hrefs(resp.Body, base)
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the URL that you want to build a sitemap for")
	flag.Parse()

	pages := get(*urlFlag)
	for _, href := range pages {
		fmt.Println(href)
	}

}
