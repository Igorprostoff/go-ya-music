package entities

type Genre struct {
	Id            string            `json:"id,omitempty"`
	Weight        int               `json:"weight,omitempty"`
	ComposerTop   bool              `json:"composerTop,omitempty"`
	Title         string            `json:"title,omitempty"`
	Titles        map[string]string `json:"titles,omitempty"`
	Images        Images            `json:"images"`
	ShowInMenu    bool              `json:"showInMenu,omitempty"`
	ShowInRegions []interface{}     `json:"showInRegions,omitempty"`
	FullTitle     string            `json:"fullTitle,omitempty"`
	UrlPart       string            `json:"urlPart,omitempty"`
	Color         string            `json:"color,omitempty"`
	RadioIcon     Icon              `json:"radioIcon"`
	SubGenres     []Genre           `json:"subGenres,omitempty"`
	HideInRegions interface{}       `json:"hideInRegions,omitempty"`
}
