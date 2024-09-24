package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(url_to_norm string) (string, error) {
	parsedUrl, err := url.Parse(url_to_norm)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %v", err)
	}

	fullPath := parsedUrl.Host + parsedUrl.Path
	fullPath = strings.ToLower(fullPath)
	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}
