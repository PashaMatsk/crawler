package main

import (
	"bytes"
	"fmt"

	"github.com/PashaMatsk/crawler/internal/fetcher"
	"github.com/PashaMatsk/crawler/internal/parser"
)

func main() {
	body, status, err := fetcher.Fetche("https://books.toscrape.com/")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	fmt.Printf("размер: %d, Статус: %d\n", len(body), status)

	lincs, err := parser.ParserLinks(bytes.NewReader(body))
	if err != nil {
		fmt.Println("Ошибка :%v", err)
	}

	for _, link := range lincs {
		fmt.Println(link)
	}
}
