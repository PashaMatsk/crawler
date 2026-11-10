package main

import (
	"fmt"
	"time"

	"github.com/PashaMatsk/crawler/internal/crawler"
)

func main() {
	start := time.Now()
	result := crawler.Crawl("https://books.toscrape.com/", 1)
	for _, p := range result {
		if p.Error != "" {
			fmt.Printf("[%d] ОШИБКА %s: %s\n", p.Depth, p.URL, p.Error)
			continue
		}
		fmt.Printf("[%d] %d %s (ссылок: %d)\n", p.Depth, p.Status, p.URL, len(p.Links))
	}
	fmt.Println("всего страниц:", len(result))
	fmt.Println("Время: ", time.Since(start))
}
