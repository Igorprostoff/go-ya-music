package entities

import (
	"encoding/json"
	"io"
)

type Library struct {
	PlaylistUuid string `json:"playlistUuid"`
	Trackslist
}

type LibraryResult struct {
	Result struct {
		Library Library `json:"library"`
	}
	ResponseWithoutResult
}

func (s *LibraryResult) ParseFromReader(i io.Reader) error {
	b, err := io.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, s)
}
