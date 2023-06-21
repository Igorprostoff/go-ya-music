package entities

import (
	"encoding/json"
	"io"
)

type Trackslist struct {
	Uid      int     `json:"uid"`
	Revision int     `json:"revision"`
	Tracks   []Track `json:"tracks"`
}

type TrackslistResult struct {
	ResponseWithoutResult
	Result Trackslist `json:"result"`
}

func (s *TrackslistResult) ParseFromReader(i io.Reader) error {
	b, err := io.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, s)
}
