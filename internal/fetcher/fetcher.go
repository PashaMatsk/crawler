package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 10 * time.Second}

func Fetche(url string) ([]byte, int, error) {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Запрос: %s: %v\n", url, err)
		return nil, 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, 0, fmt.Errorf("Плохой статус: %d\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("Чтение тела: %w\n", err)
	}

	return body, resp.StatusCode, nil
}
