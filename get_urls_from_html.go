package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody string, url *url.URL) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	rootNode, err := html.Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %v", err)
	}

	var urls []string
	err = traverseNodes(rootNode, url, &urls)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func traverseNodes(node *html.Node, url *url.URL, found *[]string) error {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				anchor, err := url.Parse(a.Val)
				if err != nil {
					fmt.Printf("couldn't parse href '%v': %v\n", a.Val, err)
					continue
				}
				*found = append(*found, url.ResolveReference(anchor).String())
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if err := traverseNodes(c, url, found); err != nil {
			return err
		}
	}

	return nil
}
