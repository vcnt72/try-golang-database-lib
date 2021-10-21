package request

type PaginateRequest struct {
	Page    uint `json:"page"`
	PerPage uint `json:"limit"`
}
