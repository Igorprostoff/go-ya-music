package feed

import (
	"github.com/Igorprostoff/go-ya-music/album"
	"github.com/Igorprostoff/go-ya-music/track"
)

type Album_event struct {
	Album  album.Album
	Tracks track.Track
}
