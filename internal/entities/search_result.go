package entities

type SearchResult struct {
	Type    string        `json:"type,omitempty"`
	Total   int           `json:"total,omitempty"`
	PerPage int           `json:"perPage,omitempty"`
	Order   int           `json:"order,omitempty"`
	Results []interface{} `json:"results,omitempty"` //List[T]
}
