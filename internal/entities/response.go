package entities

import (
	"encoding/json"
	"io"
)

type ResponseWithoutResult struct {
	Data             map[string]interface{} `json:"data,omitempty"`
	InvocationInfo   InvocationInfo         `json:"invocationInfo,omitempty"`
	Error            string                 `json:"error,omitempty"`
	ErrorDescription string                 `json:"errorDescription,omitempty"`
}

func (s *ResponseWithoutResult) ParseFromReader(i io.Reader) error {
	b, err := io.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, s)
}
func (r *ResponseWithoutResult) ParseFromBody(body io.ReadCloser) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyBytes, r)
}
