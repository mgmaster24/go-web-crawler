package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("got Network error: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return "", fmt.Errorf(
			"got HTTP error status code:%d",
			res.StatusCode,
		)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("got not-HTML response:%s", contentType)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %v", err)
	}

	return string(bytes), nil
}
