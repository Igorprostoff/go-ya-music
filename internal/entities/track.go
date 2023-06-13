package entities

type Track struct {
	Id                             interface{}        `json:"id,omitempty"` //Union[str, int]
	Title                          string             `json:"title,omitempty"`
	Available                      bool               `json:"available,omitempty"`
	Artists                        []Artist           `json:"artists,omitempty"`
	Albums                         []Album            `json:"albums,omitempty"`
	AvailableForPremiumUsers       bool               `json:"availableForPremiumUsers,omitempty"`
	LyricsAvailable                bool               `json:"lyricsAvailable,omitempty"`
	PoetryLoverMatches             []PoetryLoverMatch `json:"poetryLoverMatches,omitempty"`
	Best                           bool               `json:"best,omitempty"`
	RealId                         interface{}        `json:"realId,omitempty"` //Union[str, int]
	OgImage                        string             `json:"ogImage,omitempty"`
	Type                           string             `json:"type,omitempty"`
	CoverUri                       string             `json:"coverUri,omitempty"`
	Major                          Major              `json:"major"`
	DurationMs                     int                `json:"durationMs,omitempty"`
	StorageDir                     string             `json:"storageDir,omitempty"`
	FileSize                       int                `json:"fileSize,omitempty"`
	Substituted                    *Track             `json:"substituted,omitempty"`
	MatchedTrack                   *Track             `json:"matchedTrack,omitempty"`
	Normalization                  Normalization      `json:"normalization"`
	Error                          string             `json:"error,omitempty"`
	CanPublish                     bool               `json:"canPublish,omitempty"`
	State                          string             `json:"state,omitempty"`
	DesiredVisibility              string             `json:"desiredVisibility,omitempty"`
	Filename                       string             `json:"filename,omitempty"`
	UserInfo                       User               `json:"userInfo"`
	MetaData                       MetaData           `json:"metaData"`
	Regions                        []string           `json:"regions,omitempty"`
	AvailableAsRbt                 bool               `json:"availableAsRbt,omitempty"`
	ContentWarning                 string             `json:"contentWarning,omitempty"`
	Explicit                       bool               `json:"explicit,omitempty"`
	PreviewDurationMs              int                `json:"previewDurationMs,omitempty"`
	AvailableFullWithoutPermission bool               `json:"availableFullWithoutPermission,omitempty"`
	Version                        string             `json:"version,omitempty"`
	RememberPosition               bool               `json:"rememberPosition,omitempty"`
	BackgroundVideoUri             string             `json:"backgroundVideoUri,omitempty"`
	ShortDescription               string             `json:"shortDescription,omitempty"`
	IsSuitableForChildren          bool               `json:"isSuitableForChildren,omitempty"`
	TrackSource                    string             `json:"trackSource,omitempty"`
	AvailableForOptions            []string           `json:"availableForOptions,omitempty"`
	R128                           R128               `json:"r128"`
	LyricsInfo                     LyricsInfo         `json:"lyricsInfo"`
	TrackSharingFlag               string             `json:"trackSharingFlag,omitempty"`
}
