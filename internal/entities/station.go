package entities

type Station struct {
	Id              Id           `json:"id"`
	Name            string       `json:"name,omitempty"`
	Icon            Icon         `json:"icon"`
	MtsIcon         Icon         `json:"mtsIcon"`
	GeocellIcon     Icon         `json:"geocellIcon"`
	IdForFrom       string       `json:"idForFrom,omitempty"`
	Restrictions    Restrictions `json:"restrictions"`
	Restrictions2   Restrictions `json:"restrictions2"`
	FullImageUrl    string       `json:"fullImageUrl,omitempty"`
	MtsFullImageUrl string       `json:"mtsFullImageUrl,omitempty"`
	ParentId        Id           `json:"parentId"`
}
