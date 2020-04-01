package model

import "gomodtest/crawler/engine"

type SearchResult struct {
	Hits     int64
	Start    int
	Query    string
	PrevFrom int
	NextFrom int
	Items    []engine.Item
}
