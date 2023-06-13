package entities

type TrackShort struct {
	Id            int    `json:"id,omitempty"`
	Timestamp     string `json:"timestamp,omitempty"`
	AlbumId       string `json:"albumId,omitempty"`
	PlayCount     int    `json:"playCount,omitempty"`
	Recent        bool   `json:"recent,omitempty"`
	Chart         Chart  `json:"chart"`
	Track         Track  `json:"track"`
	OriginalIndex int    `json:"originalIndex,omitempty"`
}
