package crawler

import (
	"bytes"

	"github.com/PashaMatsk/crawler/internal/fetcher"
	"github.com/PashaMatsk/crawler/internal/parser"
)

type PageResult struct {
	URL    string
	Status int
	Depth  int
	Links  []string
	Error  string
}

func Crawl(startURL string, maxDepth int) []PageResult {
	placedlinks := make(map[string]bool)
	placedlinks[startURL] = true
	level := []string{startURL}
	var result []PageResult

	for i := 0; i <= maxDepth; i++ {
		var next []string
		for _, u := range level {
			body, status, err := fetcher.Fetche(u)
			if err != nil {
				result = append(result, PageResult{URL: u, Status: status, Depth: i, Error: err.Error()})
				continue
			}

			links, err := parser.ParserLinks(bytes.NewReader(body))
			if err != nil {
				result = append(result, PageResult{URL: u, Status: status, Depth: i, Error: err.Error()})
				continue
			}

			links, err = parser.NormalizeUrl(u, links)
			if err != nil {
				result = append(result, PageResult{URL: u, Status: status, Depth: i, Error: err.Error()})
				continue
			}
			result = append(result, PageResult{URL: u, Status: status, Depth: i, Links: links})

			for _, s := range links {
				if placedlinks[s] {
					continue
				}
				placedlinks[s] = true
				next = append(next, s)
			}

		}

		level = next
	}
	return result
}
