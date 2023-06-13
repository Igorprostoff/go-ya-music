package entities

import (
	"encoding/json"
	"io"
)

type ResponseWithoutResult struct {
	Data             map[string]interface{} `json:"data"`
	InvocationInfo   InvocationInfo         `json:"invocationInfo"`
	Error            string                 `json:"error"`
	ErrorDescription string                 `json:"errorDescription"`
}

func (r *ResponseWithoutResult) ParseFromBody(body io.ReadCloser) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyBytes, r)
}
