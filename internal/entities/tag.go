package entities

type Tag struct {
	Id            string `json:"id,omitempty"`
	Value         string `json:"value,omitempty"`
	Name          string `json:"name,omitempty"`
	OgDescription string `json:"ogDescription,omitempty"`
	OgImage       string `json:"ogImage,omitempty"`
}
