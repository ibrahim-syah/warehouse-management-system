package dto

import "math"

type Response struct {
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Error     interface{} `json:"error,omitempty"`
	Paginator *Paginator  `json:"paginator,omitempty"`
}

type Paginator struct {
	CurrentPage  int `json:"current_page"`
	PerPage      int `json:"limit_per_page"`
	PreviousPage int `json:"back_page"`
	NextPage     int `json:"next_page"`
	TotalRecords int `json:"total_records"`
	TotalPages   int `json:"total_pages"`
}

func MappingPaginator(page, limit, totalAllRecords int) Paginator {
	var totalPage int
	if limit > 0 {
		totalPage = int(math.Ceil(float64(totalAllRecords) / float64(limit)))
	}
	prev := page
	next := page
	if page != 1 {
		prev = page - 1
	}

	if page != totalPage {
		next = page + 1
	}

	p := Paginator{
		CurrentPage:  page,
		PerPage:      limit,
		PreviousPage: prev,
		NextPage:     next,
		TotalRecords: totalAllRecords,
		TotalPages:   totalPage,
	}

	return p
}
