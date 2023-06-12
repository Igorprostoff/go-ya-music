package artist

import (
	"github.com/Igorprostoff/go-ya-music/album"
	"github.com/Igorprostoff/go-ya-music/pager"
)

type Artist_albums struct {
	Albums []album.Album
	Pager  pager.Pager
}
