package internal

type (
	PaginationRequest struct {
		Search  string `form:"search"`
		Page    int    `form:"page"`
		PerPage int    `form:"perPage"`
	}

	PaginationResponse struct {
		Page    int   `json:"page"`
		PerPage int   `json:"perPage"`
		MaxPage int64 `json:"maxPage"`
		Count   int64 `json:"count"`
	}
)

func (p *PaginationRequest) GetOffset() int {
	return (p.Page - 1) * p.PerPage
}

func (p *PaginationResponse) GetLimit() int {
	return p.PerPage
}

func (p *PaginationResponse) GetPage() int {
	return p.Page
}
