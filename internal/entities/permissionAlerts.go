package entities

import (
	"encoding/json"
	"io"
)

type PermissionAlerts struct {
	Alerts []string `json:"alerts"`
}

type PermissionAlertsResult struct {
	ResponseWithoutResult
	Result PermissionAlerts
}

func (s *PermissionAlertsResult) ParseFromReader(i io.Reader) error {
	b, err := io.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, s)
}
