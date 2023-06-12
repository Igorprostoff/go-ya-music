package track

import (
	"github.com/Igorprostoff/go-ya-music/album"
	"github.com/Igorprostoff/go-ya-music/artist"
	"github.com/Igorprostoff/go-ya-music/playlist"
)

type Track struct {
	Id                                interface{} //Union[str, int]
	Title                             string
	Available                         bool
	Artists                           []artist.Artist
	Albums                            []album.Album
	Available_for_premium_users       bool
	Lyrics_available                  bool
	Poetry_lover_matches              []Poetry_lover_match
	Best                              bool
	Real_id                           interface{} //Union[str, int]
	Og_image                          string
	Type                              string
	Cover_uri                         string
	Major                             Major
	Duration_ms                       int
	Storage_dir                       string
	File_size                         int
	Substituted                       *Track
	Matched_track                     *Track
	Normalization                     Normalization
	Error                             string
	Can_publish                       bool
	State                             string
	Desired_visibility                string
	Filename                          string
	User_info                         playlist.User
	Meta_data                         Meta_data
	Regions                           []string
	Available_as_rbt                  bool
	Content_warning                   string
	Explicit                          bool
	Preview_duration_ms               int
	Available_full_without_permission bool
	Version                           string
	Remember_position                 bool
	Background_video_uri              string
	Short_description                 string
	Is_suitable_for_children          bool
	Track_source                      string
	Available_for_options             []string
	R128                              R128
	Lyrics_info                       Lyrics_info
	Track_sharing_flag                string
}
