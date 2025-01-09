package repository

import "math"

// Pagination struct for response client
type Pagination struct {
	CurrentPage int   `json:"current_page,omitempty"`
	PageSize    int   `json:"page_size,omitempty"`
	TotalPage   int   `json:"total_page, omitempty"`
	TotalResult int   `json:"total_result,omitempty"`
	Pages       []int `json:"pages,omitempty"`
}

// Paginate data
func Paginate(result int, page int, limit int) interface{} {
	var paginate Pagination
	var pages []int
	var maxPages, totalPages int

	maxPages = 10
	totalPages = int(math.Ceil(float64(result) / float64(limit)))

	if page < 1 {
		page = 1
	} else if page > totalPages {
		return nil
	}

	var startPage, endPage int
	if totalPages <= maxPages {
		// total pages less than max -> show all pages
		startPage = 1
		endPage = totalPages
	} else {
		var maxPagesBeforeCurrentPage = math.Floor(float64(maxPages) / 2)
		var maxPagesAfterCurrentPage = math.Floor(float64(maxPages)/2) - 1

		if float64(page) <= maxPagesBeforeCurrentPage {
			// current page near the start
			startPage = 1
			endPage = maxPages
		} else if float64(page)+maxPagesAfterCurrentPage >= float64(totalPages) {
			// current page near the end
			startPage = totalPages - maxPages + 1
			endPage = totalPages
		} else {
			// current page near the end
			startPage = int(float64(page) - maxPagesBeforeCurrentPage)
			endPage = int(float64(page) + maxPagesAfterCurrentPage)
		}
	}

	count := (endPage + 1) - startPage
	for i := 0; i < count; i++ {
		pages = append(pages, startPage+i)
	}

	paginate = Pagination{
		CurrentPage: page,
		TotalPage:   totalPages,
		TotalResult: result,
		PageSize:    limit,
		Pages:       pages,
	}

	return &paginate
}
