package feed

import "github.com/Igorprostoff/go-ya-music/track"

type Day struct {
	Day                     string
	Events                  []Event
	Tracks_to_play_with_ads []Track_with_ads
	Tracks_to_play          []track.Track
}
