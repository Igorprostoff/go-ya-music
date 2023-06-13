package entities

type Vinyl struct {
	Url       string `json:"url,omitempty"`
	Title     string `json:"title,omitempty"`
	Year      int    `json:"year,omitempty"`
	Price     int    `json:"price,omitempty"`
	Media     string `json:"media,omitempty"`
	OfferId   int    `json:"offerId,omitempty"`
	ArtistIds []int  `json:"artistIds,omitempty"`
	Picture   string `json:"picture,omitempty"`
}
