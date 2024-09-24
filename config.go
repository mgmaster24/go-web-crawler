package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func new(baseURL string) (*config, error) {
	parsedUrl, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse URL: %v", err)
	}

	return &config{
		pages:   make(map[string]int),
		baseURL: parsedUrl,
	}, nil
}
