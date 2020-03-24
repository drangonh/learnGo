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

	var limitCount = 2
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got it item %v", item)
		}
		limitCount--
		if limitCount == 0 {
			break
		}
	}
}

func Worker(r Request) (ParseResult, error) {
	//查询数据
	body, err := fetcher.Fetch(r.Url)

	log.Printf("Fetching url %s", r.Url)
	if err != nil {
		log.Printf("Fetching error url %s", r.Url)
		return ParseResult{}, err
	}

	//解析数据
	return r.ParseFunc(body), nil
}
