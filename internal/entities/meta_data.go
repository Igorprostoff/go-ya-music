package entities

type MetaData struct {
	Album    string `json:"album,omitempty"`
	Volume   int    `json:"volume,omitempty"`
	Year     int    `json:"year,omitempty"`
	Number   int    `json:"number,omitempty"`
	Genre    string `json:"genre,omitempty"`
	Lyricist string `json:"lyricist,omitempty"`
	Version  string `json:"version,omitempty"`
	Composer string `json:"composer,omitempty"`
}
