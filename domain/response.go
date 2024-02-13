package domain

type Response struct {
	Data        any    `json:"data"`
	Message     string `json:"message"`
	ElapsedTime string `json:"elapsed_time"`
}

type PagingResponse struct {
	Data      any `json:"data"`
	PerPage   int `json:"per_page"`
	Page      int `json:"page"`
	TotalData int `json:"total_data"`
	TotalPage int `json:"total_page"`
}

type FilterRequest struct {
	Page    int    `json:"page" form:"page" gorm:"default:1"`
	Limit   int    `json:"limit" form:"limit" gorm:"default:20"`
	OrderBy string `json:"order_by" form:"order_by"`
	SortBy  string `json:"sort_by" form:"sort_by"`
}
