package pagination

import (
	"errors"
	"strconv"
)

var (
	ErrPageTooLow  = errors.New("page must be greater than 0")
	ErrLimitTooLow = errors.New("limit must be greater than 0")
)

// Pagination handle with pagination params.
type Pagination struct {
	Page        int
	LimitOfPage int
}

// PaginationInit creates a new Pagination from query parameters.
func PaginationInit(pageStr, limitStr string) (*Pagination, error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	return &Pagination{Page: page, LimitOfPage: limit}, nil
}

// Offset calculates the number of items to skip.
func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.LimitOfPage
}

// Limit returns the limit for the number of items.
func (p *Pagination) Limit() int {
	return p.LimitOfPage
}

// Validate ensures the pagination parameters are within acceptable ranges.
func (p *Pagination) Validate() error {
	if p.Page < 1 {
		return ErrPageTooLow
	}
	if p.LimitOfPage < 1 {
		return ErrLimitTooLow
	}
	return nil
}
