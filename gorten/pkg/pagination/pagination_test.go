package pagination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaginationInit(t *testing.T) {
	tests := []struct {
		name          string
		pageStr       string
		limitStr      string
		expectedPage  int
		expectedLimit int
	}{
		{"Valid parameters", "2", "20", 2, 20},
		{"Invalid page, valid limit", "0", "15", 1, 15},
		{"Valid page, invalid limit", "3", "-5", 3, 10},
		{"Invalid page and limit", "-1", "0", 1, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pagination, err := PaginationInit(tt.pageStr, tt.limitStr)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.expectedPage, pagination.Page)
				assert.Equal(t, tt.expectedLimit, pagination.LimitOfPage)
			}
		})
	}
}

func TestPagination_Offset(t *testing.T) {
	tests := []struct {
		name           string
		page           int
		limit          int
		expectedOffset int
	}{
		{"Page 1", 1, 10, 0},
		{"Page 2", 2, 10, 10},
		{"Page 3", 3, 5, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pagination := &Pagination{Page: tt.page, LimitOfPage: tt.limit}
			assert.Equal(t, tt.expectedOffset, pagination.Offset())
		})
	}
}

func TestPagination_Limit(t *testing.T) {
	tests := []struct {
		name          string
		page          int
		limit         int
		expectedLimit int
	}{
		{"Limit 10", 1, 10, 10},
		{"Limit 5", 2, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pagination := &Pagination{Page: tt.page, LimitOfPage: tt.limit}
			assert.Equal(t, tt.expectedLimit, pagination.Limit())
		})
	}
}

func TestPagination_Validate(t *testing.T) {
	tests := []struct {
		name          string
		page          int
		limit         int
		expectedError error
	}{
		{"Valid parameters", 1, 10, nil},
		{"Invalid page", 0, 10, ErrPageTooLow},
		{"Invalid limit", 1, 0, ErrLimitTooLow},
		{"Invalid page and limit", 0, 0, ErrPageTooLow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pagination := &Pagination{Page: tt.page, LimitOfPage: tt.limit}
			err := pagination.Validate()
			assert.ErrorIs(t, err, tt.expectedError)
		})
	}
}
