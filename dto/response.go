package dto

type Response struct {
	Message        string          `json:"message,omitempty"`
	Data           interface{}     `json:"data,omitempty"`
	Error          interface{}     `json:"error,omitempty"`
	PaginationInfo *PaginationInfo `json:"pagination,omitempty"`
}

type PaginationInfo struct {
	CurrentPage    int `json:"current_page"`
	ItemsPerPage   int `json:"items_per_page"`
	TotalPageCount int `json:"total_page_count"`
	TotalItemCount int `json:"total_items_count"`
}
