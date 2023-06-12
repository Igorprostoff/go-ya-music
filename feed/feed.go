package feed

type Feed struct {
	Can_get_more_events bool
	Pumpkin             bool
	Is_wizard_passed    bool
	Generated_playlists []Generated_playlist
	Headlines           []interface{}
	Today               string
	Days                []Day
	Next_revision       string
}
