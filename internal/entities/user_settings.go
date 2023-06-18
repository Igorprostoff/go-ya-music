package entities

import (
	"encoding/json"
	"io"
)

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

type UserSettingsResult struct {
	ResponseWithoutResult
	Result UserSettings
}

func (s *UserSettingsResult) ParseFromReader(i io.Reader) error {
	b, err := io.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, s)
}
