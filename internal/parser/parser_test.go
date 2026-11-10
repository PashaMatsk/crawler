package parser

import (
	"slices"
	"strings"
	"testing"
)

func TestParselinks(t *testing.T) {
	page := `<html>
	<body>
		<p>Привет <a href="/about">о нас</a></p>
		<a href="/books">книги</a>
		<a>ссылка без href</a>
		<p>просто текст</p>
	</body>
</html>`
	links, err := ParserLinks(strings.NewReader(page))
	if err != nil {
		t.Fatalf("неожиданная ошибка: %v", err)
	}

	want := []string{"/about", "/books"}
	if !slices.Equal(links, want) {
		t.Fatalf("ждали %v, получили %v", want, links)
	}
}
