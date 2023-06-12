package search

import (
	"github.com/Igorprostoff/go-ya-music/album"
	"github.com/Igorprostoff/go-ya-music/artist"
	"github.com/Igorprostoff/go-ya-music/playlist"
	"github.com/Igorprostoff/go-ya-music/track"
	"github.com/Igorprostoff/go-ya-music/video"
)

type Search struct {
	Search_request_id  string
	Text               string
	Best               Best
	Albums             []album.Album
	Artists            []artist.Artist
	Playlists          []playlist.Playlist
	Tracks             []track.Track
	Videos             []video.Video
	Users              []playlist.User
	Podcasts           []album.Album
	Podcast_episodes   []track.Track
	Type               string
	Page               int
	Per_page           int
	Misspell_result    string
	Misspell_original  string
	Misspell_corrected bool
	Nocorrect          bool
}
