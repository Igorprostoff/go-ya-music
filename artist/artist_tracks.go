package artist

import (
	"github.com/Igorprostoff/go-ya-music/pager"
	"github.com/Igorprostoff/go-ya-music/track"
)

type Artist_tracks struct {
	Tracks track.Track
	Pager  pager.Pager
}
