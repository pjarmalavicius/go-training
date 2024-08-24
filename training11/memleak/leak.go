package memleak

import (
	"fmt"
	"net/http"
)

func CallHTTP() error {
	resp, err := http.Get("https://example.com")
	if err != nil {
		return fmt.Errorf("failed to get: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
