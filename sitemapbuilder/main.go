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

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func hrefs(body io.Reader, base string) []string {
	var ret []string
	links, _ := linkparser.Parse(body)
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}

	return ret
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
	// fmt.Println("Request URL: ", reqURL)
	// fmt.Println("Base URL: ", base)

	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}
	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		for url, _ := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}
			for _, link := range get(url) {
				nq[link] = struct{}{}
			}
		}
	}

	ret := make([]string, 0, len(seen))
	for url, _ := range seen {
		ret = append(ret, url)
	}

	return ret
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the URL that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 3, "the maximum depth to traverse")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)
	for _, href := range pages {
		fmt.Println(href)
	}

}
