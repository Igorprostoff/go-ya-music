package entities

type Artist struct {
	Id                    int
	Error                 string
	Reason                string
	Name                  string
	Cover                 Cover
	Various               bool
	Composer              bool
	Genres                []string
	OgImage               string
	OpImage               string
	NoPicturesFromSearch  interface{}
	Counts                Counts
	Available             bool
	Ratings               Raitings
	Links                 []Link
	TicketsAvailable      bool
	LikesCount            int
	PopularTracks         []Track
	Regions               []string
	Decomposed          []string
	FullNames           interface{}
	HandMadeDescription string
	Description         Description
	Countries       []string
	EnWikipediaLink string
	DbAliases       []string
	Aliases           interface{}
	InitDate string
	EndDate  string
	YaMoneyId string
}
