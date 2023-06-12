package search

type Best struct {
	Type   string
	Result interface{} //TODO: [Union[Track, Artist, Album, Playlist, Video]]
	Text   string
}
