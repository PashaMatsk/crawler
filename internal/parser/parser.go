package parser

import (
	"io"

	"golang.org/x/net/html"
)

func ParserLinks(body io.Reader) ([]string, error) {
	links := make([]string, 0, 7)
	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}
	links = walktree(doc, links)
	return links, nil
}

func walktree(n *html.Node, links []string) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, atr := range n.Attr {
			if atr.Key == "href" {
				links = append(links, atr.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = walktree(c, links)
	}
	return links
}
