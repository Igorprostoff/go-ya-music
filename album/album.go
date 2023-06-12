package album

import (
	"github.com/Igorprostoff/go-ya-music/artist"
	"github.com/Igorprostoff/go-ya-music/track"
)

type Album struct {
	Id                          int
	Error                       string
	Title                       string
	Track_count                 int
	Artists                     []artist.Artist
	Labels                      []string
	Available                   bool
	Available_for_premium_users bool
	Version                     string
	Cover_uri                   string
	Content_warning             string
	Original_release_year       interface{}
	Genre                       string
	Text_color                  string
	Short_description           string
	Description                 string
	Is_premiere                 bool
	Is_banner                   bool
	Meta_type                   string
	Storage_dir                 string
	Og_image                    string
	Buy                         []interface{}
	Recent                      bool
	Very_important              bool
	Available_for_mobile        bool
	Available_partially         bool
	Bests                       []int
	Duplicates                  []Album
	Prerolls                    []interface{}
	Volumes                     [][]track.Track
	Year                        int
	Release_date                string
	Type                        string
	Track_position              Track_position
	Regions                     []string
	Available_as_rbt            bool
	Lyrics_available            bool
	Remember_position           bool
	Albums                      []Album
	Duration_ms                 int
	Explicit                    bool
	Start_date                  string
	Likes_count                 int
	Deprecation                 Deprecation
	Available_regions           []string
	Available_for_options       []string
}
