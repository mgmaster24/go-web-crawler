package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseUrl := args[0]
	fmt.Println("starting crawl of:", baseUrl)
  config, err := configure(baseUrl, 5)
  if err != nil {
    fmt.Println("failed to create config")
    os.Exit(1)
  }

  config.wg.Add(1)
	go config.crawlPage(baseUrl)
  config.wg.Wait()

	for normalizedURL, count := range config.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
