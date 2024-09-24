package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedCurrURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normailizeURL: %v\n", err)
		return
	}

	if _, visited := cfg.pages[normalizedCurrURL]; visited {
		cfg.pages[normalizedCurrURL]++
		return
	}

	cfg.pages[normalizedCurrURL] = 1

	fmt.Printf("crawling %s\n", rawCurrentURL)
	currentRawHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - getHTML: %v", err)
		return
	}

	urls, err := getURLsFromHTML(currentRawHTML, cfg.baseURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v\n", err)
		return
	}

	for _, url := range urls {
		cfg.crawlPage(url)
	}
}
