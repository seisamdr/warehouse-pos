package pagination

import "math"

type PaginationResponse struct {
	CurrentPage  int  `json:"current_page"`
	TotalPages   int  `json:"total_pages"`
	TotalRecords int  `json:"total_records"`
	Limit        int  `json:"limit"`
	HasNext      bool `json:"has_next"`
	HasPrev      bool `json:"has_prev"`
}

func CalculatePagination(page, limit, totalRecords int) PaginationResponse {
	totalPage := int(math.Ceil(float64(totalRecords) / float64(limit)))
	if totalPage == 0 {
		totalPage = 1
	}

	return PaginationResponse{
		CurrentPage:  page,
		TotalPages:   totalPage,
		TotalRecords: totalRecords,
		Limit:        limit,
		HasNext:      page < totalPage,
		HasPrev:      page > 1,
	}
}