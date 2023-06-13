package entities

type UserSettings struct {
	Uid                       int    `json:"uid,omitempty"`
	LastFmScrobblingEnabled   bool   `json:"lastFmScrobblingEnabled,omitempty"`
	ShuffleEnabled            bool   `json:"shuffleEnabled,omitempty"`
	VolumePercents            int    `json:"volumePercents,omitempty"`
	Modified                  string `json:"modified,omitempty"`
	FacebookScrobblingEnabled bool   `json:"facebookScrobblingEnabled,omitempty"`
	AddNewTrackOnPlaylistTop  bool   `json:"addNewTrackOnPlaylistTop,omitempty"`
	UserMusicVisibility       string `json:"userMusicVisibility,omitempty"`
	UserSocialVisibility      string `json:"userSocialVisibility,omitempty"`
	RbtDisabled               bool   `json:"rbtDisabled,omitempty"`
	Theme                     string `json:"theme,omitempty"`
	PromosDisabled            bool   `json:"promosDisabled,omitempty"`
	AutoPlayRadio             bool   `json:"autoPlayRadio,omitempty"`
	SyncQueueEnabled          bool   `json:"syncQueueEnabled,omitempty"`
	AdsDisabled               bool   `json:"adsDisabled,omitempty"`
	DiskEnabled               bool   `json:"diskEnabled,omitempty"`
	ShowDiskTracksInLibrary   bool   `json:"showDiskTracksInLibrary,omitempty"`
}
