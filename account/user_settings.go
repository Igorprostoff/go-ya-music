package account

type User_settings struct {
	Uid                           int
	Last_fm_scrobbling_enabled    bool
	Shuffle_enabled               bool
	Volume_percents               int
	Modified                      string
	Facebook_scrobbling_enabled   bool
	Add_new_track_on_playlist_top bool
	User_music_visibility         string
	User_social_visibility        string
	Rbt_disabled                  bool
	Theme                         string
	Promos_disabled               bool
	Auto_play_radio               bool
	Sync_queue_enabled            bool
	Ads_disabled                  bool
	Disk_enabled                  bool
	Show_disk_tracks_in_library   bool
}
