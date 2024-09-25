package main

import (
	"fmt"
	"sort"
)

type siteVisit struct {
	visits int
	site   string
}

func print_report(baseURL string, pages map[string]int) {
	fmt.Println("=========================================================")
	fmt.Printf("\tREPORT for %s\n", baseURL)
	fmt.Println("=========================================================")

	siteVisits := []siteVisit{}
	for k, v := range pages {
		siteVisits = append(siteVisits, siteVisit{
			site:   k,
			visits: v,
		})
	}

	sort.Slice(siteVisits, func(i, j int) bool {
		return siteVisits[i].visits > siteVisits[j].visits
	})

	for _, visit := range siteVisits {
		fmt.Printf("Found %d internal links to %s\n", visit.visits, visit.site)
	}
}
