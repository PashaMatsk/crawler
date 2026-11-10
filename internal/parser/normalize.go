package parser

import (
	"fmt"
	"net/url"
)

func NormalizeUrl(pageURL string, links []string) ([]string, error) {
	base, err := url.Parse(pageURL)
	if err != nil {
		return nil, fmt.Errorf("Не удалось пропарсить ссылку: %v", err)
	}
	//объеденение строк
	var result []string
	seen := make(map[string]bool)

	for _, link := range links {
		ref, err := url.Parse(link)
		if err != nil {
			continue
		}
		//удаление якоря и склейка ссылки
		abs := base.ResolveReference(ref)
		abs.Fragment = ""

		if abs.Scheme != "http" && abs.Scheme != "https" {
			continue
		}

		if abs.Host != base.Host {
			continue
		}
		//работа с картой
		s := abs.String()
		if seen[s] {
			continue
		}
		seen[s] = true
		result = append(result, s)
	}

	return result, nil

}
