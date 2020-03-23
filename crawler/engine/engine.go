package engine

import (
	"gomodtest/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		body, err := fetcher.Fetch(r.Url)

		log.Printf("Fetching url %s", r.Url)
		if err != nil {
			log.Printf("Fetching error url %s", r.Url)
			continue
		}

		parseResult := r.ParseFunc(body)

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got it item %v", item)
		}
	}
}
