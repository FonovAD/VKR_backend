package pagination

// PaginationParams represents pagination parameters
type PaginationParams struct {
	Page     int `json:"page" query:"page"`         // Page number (1-based)
	PageSize int `json:"page_size" query:"page_size"` // Number of items per page
}

// PaginatedResponse represents a paginated response
type PaginatedResponse[T any] struct {
	Data       []T   `json:"data"`
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalCount int64 `json:"total_count"`
	TotalPages int   `json:"total_pages"`
}

// ValidateAndSetDefaults validates pagination parameters and sets defaults
func (p *PaginationParams) ValidateAndSetDefaults() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10 // Default page size
	}
	if p.PageSize > 100 {
		p.PageSize = 100 // Max page size
	}
}

// Offset calculates the offset for SQL queries
func (p *PaginationParams) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// Limit returns the limit for SQL queries
func (p *PaginationParams) Limit() int {
	return p.PageSize
}

// CalculateTotalPages calculates total pages based on total count
func CalculateTotalPages(totalCount int64, pageSize int) int {
	if pageSize <= 0 {
		return 0
	}
	totalPages := int(totalCount) / pageSize
	if int(totalCount)%pageSize > 0 {
		totalPages++
	}
	return totalPages
}

