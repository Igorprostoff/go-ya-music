package entities

type Playlist struct {
	Owner                 User            `json:"owner"`
	Cover                 Cover           `json:"cover"`
	MadeFor               MadeFor         `json:"madeFor"`
	PlayCounter           PlayCounter     `json:"playCounter"`
	PlaylistAbsence       PlaylistAbsence `json:"playlistAbsence"`
	Uid                   int             `json:"uid,omitempty"`
	Kind                  int             `json:"kind,omitempty"`
	Title                 string          `json:"title,omitempty"`
	TrackCount            int             `json:"trackCount,omitempty"`
	Tags                  []interface{}   `json:"tags,omitempty"`
	Revision              int             `json:"revision,omitempty"`
	Snapshot              int             `json:"snapshot,omitempty"`
	Visibility            string          `json:"visibility,omitempty"`
	Collective            bool            `json:"collective,omitempty"`
	UrlPart               string          `json:"urlPart,omitempty"`
	Created               string          `json:"created,omitempty"`
	Modified              string          `json:"modified,omitempty"`
	Available             bool            `json:"available,omitempty"`
	IsBanner              bool            `json:"isBanner,omitempty"`
	IsPremiere            bool            `json:"isPremiere,omitempty"`
	DurationMs            int             `json:"durationMs,omitempty"`
	OgImage               string          `json:"ogImage,omitempty"`
	OgTitle               string          `json:"ogTitle,omitempty"`
	OgDescription         string          `json:"ogDescription,omitempty"`
	Image                 string          `json:"image,omitempty"`
	CoverWithoutText      Cover           `json:"coverWithoutText"`
	Contest               Contest         `json:"contest"`
	BackgroundColor       string          `json:"backgroundColor,omitempty"`
	TextColor             string          `json:"textColor,omitempty"`
	IdForFrom             string          `json:"idForFrom,omitempty"`
	DummyDescription      string          `json:"dummyDescription,omitempty"`
	DummyPageDescription  string          `json:"dummyPageDescription,omitempty"`
	DummyCover            Cover           `json:"dummyCover"`
	DummyRolloverCover    Cover           `json:"dummyRolloverCover"`
	OgData                OpenGraphData   `json:"ogData"`
	Branding              Brand           `json:"branding"`
	MetrikaId             int             `json:"metrikaId,omitempty"`
	Coauthors             []int           `json:"coauthors,omitempty"`
	TopArtist             []Artist        `json:"topArtist,omitempty"`
	RecentTracks          []TrackId       `json:"recentTracks,omitempty"`
	Tracks                []TrackShort    `json:"tracks,omitempty"`
	Prerolls              []interface{}   `json:"prerolls,omitempty"`
	LikesCount            int             `json:"likesCount,omitempty"`
	SimilarPlaylists      []Playlist      `json:"similarPlaylists,omitempty"`
	LastOwnerPlaylists    []Playlist      `json:"lastOwnerPlaylists,omitempty"`
	GeneratedPlaylistType string          `json:"generatedPlaylistType,omitempty"`
	AnimatedCoverUri      string          `json:"animatedCoverUri,omitempty"`
	EverPlayed            bool            `json:"everPlayed,omitempty"`
	Description           string          `json:"description,omitempty"`
	DescriptionFormatted  string          `json:"descriptionFormatted,omitempty"`
	PlaylistUuid          string          `json:"playlistUuid,omitempty"`
	Type                  string          `json:"type,omitempty"`
	Ready                 bool            `json:"ready,omitempty"`
	IsForFrom             interface{}     `json:"isForFrom,omitempty"`
	Regions               interface{}     `json:"regions,omitempty"`
	CustomWave            CustomWave      `json:"customWave"`
	Pager                 Pager           `json:"pager"`
}
