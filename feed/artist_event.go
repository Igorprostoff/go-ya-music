package feed

type Artist_event struct {
	Artist                          Artist
	Tracks                          []Track
	Similar_to_artists_from_history []Artist
	Subscribed                      bool
}
