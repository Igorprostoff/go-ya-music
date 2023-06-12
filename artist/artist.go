package artist

import (
	"github.com/Igorprostoff/go-ya-music/album"
	"github.com/Igorprostoff/go-ya-music/track"
)

type Artist struct {
	Id                      int
	Error                   string
	Reason                  string
	Name                    string
	Cover                   album.Cover
	Various                 bool
	Composer                bool
	Genres                  []string
	Og_image                string
	Op_image                string
	No_pictures_from_search interface{}
	Counts                  Counts
	Available               bool
	Ratings                 Raitings
	Links                   []Link
	Tickets_available       bool
	Likes_count             int
	Popular_tracks          []track.Track
	Regions                 []string
	Decomposed              []string
	Full_names              interface{}
	Hand_made_description   string
	Description             Description
	Countries               []string
	En_wikipedia_link       string
	Db_aliases              []string
	Aliases                 interface{}
	Init_date               string
	End_date                string
	Ya_money_id             string
}
