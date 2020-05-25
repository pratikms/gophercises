package linkparser

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="">...</a>) in an HTML document
type Link struct {
	Href string
	Text string
}

func buildLinks(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = "TODO: Parse the text..."
	return ret
}

func linkNodes(n *html.Node) []*html.Node {
	var ret []*html.Node
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

// Parse will take an HTML document and will return a slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var links []Link
	nodes := linkNodes(doc)
	for _, node := range nodes {
		links = append(links, buildLinks(node))
		fmt.Println(node)
	}
	return links, nil
}
