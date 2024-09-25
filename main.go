package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println("not enough arguments provided")
		os.Exit(1)
	}
	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseUrl := args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("failed to parse argument to int")
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("failed to parse argument to int")
		os.Exit(1)
	}

	fmt.Println("starting crawl of:", baseUrl)
	config, err := configure(baseUrl, maxConcurrency, maxPages)
	if err != nil {
		fmt.Println("failed to create config")
		os.Exit(1)
	}

	config.wg.Add(1)
	go config.crawlPage(baseUrl)
	config.wg.Wait()

	print_report(baseUrl, config.pages)
	//for normalizedURL, count := range config.pages {
	//	fmt.Printf("%d - %s\n", count, normalizedURL)
	//}
}
