package entities

type Pager struct {
	Total   int `json:"total,omitempty"`
	Page    int `json:"page,omitempty"`
	PerPage int `json:"perPage,omitempty"`
}
