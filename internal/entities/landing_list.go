package entities

type LandingList struct {
	Type         string `json:"type,omitempty"`
	TypeForFrom  string `json:"typeForFrom,omitempty"`
	Title        string `json:"title,omitempty"`
	Id           string `json:"id,omitempty"`
	NewReleases  []int  `json:"newReleases,omitempty"`
	NewPlaylists []int  `json:"newPlaylists,omitempty"`
	Podcasts     []int  `json:"podcasts,omitempty"`
}
