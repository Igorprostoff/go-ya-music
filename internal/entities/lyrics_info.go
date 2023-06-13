package entities

type LyricsInfo struct {
	HasAvailableSyncLyrics bool `json:"hasAvailableSyncLyrics,omitempty"`
	HasAvailableTextLyrics bool `json:"hasAvailableTextLyrics,omitempty"`
}
