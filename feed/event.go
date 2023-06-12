package feed

import "github.com/Igorprostoff/go-ya-music/track"

type Event struct {
	Id            string
	Type          string
	Type_for_from string
	Title         string
	Tracks        []track.Track
	Artists       []Artist_event
	Albums        []Album_event
	Message       string
	Device        string
	Tracks_count  int
	Genre         string
}
