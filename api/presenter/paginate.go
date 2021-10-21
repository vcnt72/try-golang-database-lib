package presenter

type PaginatePresenter struct {
	Page    uint `json:"page"`
	PerPage uint `json:"per_page"`
}
