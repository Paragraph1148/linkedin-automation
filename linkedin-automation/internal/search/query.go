package search

import (
	"fmt"
	"net/url"
)

type SearchQuery struct {
	Keywords string
	Location string
	Company string
	Page int
}

func (q SearchQuery) URL() string {
	params := url.Values {}
	if q.Keywords != "" {
		params.Set("keywords", q.Keywords)
	}
	if q.Location != "" {
		params.Set("location", q.Location)
	}
	if q.Company != "" {
		params.Set("company", q.Company)
	}

	start := q.Page * 10
	return fmt.Sprintf(
		"https://www.linkedin.com/search/results/people/?%s&origin=GLOBAL_SEARCH_HEADER&start=%d",
		params.Encode(),
		start,
	)
}