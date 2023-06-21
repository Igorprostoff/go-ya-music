package entities

import (
	"encoding/json"
	"io"
)

type PromocodeStatus struct {
	Status        string
	Status_desc   string
	AccountStatus Status
}

type PromocodeStatusResult struct {
	ResponseWithoutResult
	Result PromocodeStatus `json:"result"`
}

func (s *PromocodeStatusResult) ParseFromReader(i io.Reader) error {
	b, err := io.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, s)
}
