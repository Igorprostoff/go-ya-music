package entities

type Cover struct {
	Type           string `json:"type,omitempty"`
	uri            string `json:"uri,omitempty"`
	itemsUri       string `json:"itemsUri,omitempty"`
	dir            string `json:"dir,omitempty"`
	version        string `json:"version,omitempty"`
	custom         bool   `json:"custom,omitempty"`
	isCustom       bool   `json:"isCustom,omitempty"`
	copyrightName  string `json:"copyrightName,omitempty"`
	copyrightCline string `json:"copyrightCline,omitempty"`
	prefix         string `json:"prefix,omitempty"`
	error          string `json:"error,omitempty"`
}
