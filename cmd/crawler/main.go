package main

import (
	"fmt"

	"github.com/PashaMatsk/crawler/internal/fetcher"
)

func main() {
	body, status, err := fetcher.Fetche("https://books.toscrape.com/")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}
	fmt.Printf("размер: %d, Статус: %d\n", len(body), status)
}
