package entities

type ChartInfoMenuItem struct {
	Title    string `json:"title,omitempty"`
	Url      string `json:"url,omitempty"`
	Selected bool   `json:"selected,omitempty"`
}
