package entities

type InvocationInfo struct {
	Hostname           string `json:"hostname,omitempty"`
	ReqId              string `json:"reqId,omitempty"`
	ExecDurationMillis int    `json:"execDurationMillis,omitempty"`
}
