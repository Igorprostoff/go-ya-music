package artist

import (
	"github.com/Igorprostoff/go-ya-music/album"
	"github.com/Igorprostoff/go-ya-music/landing"
	"github.com/Igorprostoff/go-ya-music/playlist"
	"github.com/Igorprostoff/go-ya-music/track"
	"github.com/Igorprostoff/go-ya-music/video"
)

type Brief_info struct {
	Artist           Artist
	Albums           []album.Album
	Playlists        []playlist.Playlist
	Also_albums      []album.Album
	Last_release_ids []int
	Last_releases    []album.Album
	Popular_tracks   []track.Track
	Similar_artists  []Artist
	All_covers       []album.Cover
	Concerts         interface{}
	Videos           []video.Video
	Vinyls           []Vinyl
	Has_promotions   bool
	Playlist_ids     []playlist.Playlist_id
	Tracks_in_chart  []landing.Chart
}
