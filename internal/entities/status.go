package entities

import (
	"encoding/json"
	"io"
)

type Status struct {
	Account        Account      `json:"account"`
	Permissions    Permissions  `json:"permissions"`
	Advertisement  string       `json:"advertisement,omitempty"`
	Subscription   Subscription `json:"subscription"`
	CacheLimit     int          `json:"cacheLimit,omitempty"`
	Subeditor      bool         `json:"subeditor,omitempty"`
	SubeditorLevel int          `json:"subeditorLevel,omitempty"`
	Plus           Plus         `json:"plus"`
	DefaultEmail   string       `json:"defaultEmail,omitempty"`
	SkipsPerHour   int          `json:"skipsPerHour,omitempty"`
	StationExists  bool         `json:"stationExists,omitempty"`
	StationData    StationData  `json:"stationData"`
	BarBelow       Alert        `json:"barBelow"`
	PremiumRegion  int          `json:"premiumRegion,omitempty"`
	Experiment     int          `json:"experiment,omitempty"`
	PretrialActive bool         `json:"pretrialActive,omitempty"`
	Userhash       string       `json:"userhash,omitempty"`
}

type StatusResult struct {
	ResponseWithoutResult
	Result Status
}

func (s *StatusResult) ParseFromBytes(b []byte) error {
	return json.Unmarshal(b, s)
}

func (s *StatusResult) ParseFromReader(i io.Reader) error {
	b, err := io.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, s)
}
