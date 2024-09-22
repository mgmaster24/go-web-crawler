package main

import (
	"fmt"
	"net/url"
)

func NormalizeURL(url_to_norm string) (string, error) {
	parsedUrl, err := url.Parse(url_to_norm)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s", parsedUrl.Hostname(), parsedUrl.Path), nil
}
