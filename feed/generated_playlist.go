package feed

type Generated_playlist struct {
	Type                string
	Ready               bool
	Notify              bool
	Data                Playlist
	Description         []interface{}
	Preview_description string
}
