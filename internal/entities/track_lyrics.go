package entities

type TrackLyrics struct {
	DownloadUrl     string      `json:"downloadUrl,omitempty"`
	LyricId         int         `json:"lyricId,omitempty"`
	ExternalLyricId string      `json:"externalLyricId,omitempty"`
	Writers         []string    `json:"writers,omitempty"`
	Major           LyricsMajor `json:"major"`
}
