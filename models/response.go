package models

type ErrorResp struct {
	Error string `json:"error"`
	Details interface{} `json:"details,omitempty"`

}

