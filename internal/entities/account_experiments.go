package entities

import (
	"encoding/json"
	"io"
)

type AccountExperiments map[string]interface{}

type AccountExperimentsResult struct {
	ResponseWithoutResult
	Result AccountExperiments `json:"result"`
}

func (s *AccountExperimentsResult) ParseFromReader(i io.Reader) error {
	b, err := io.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, s)
}
