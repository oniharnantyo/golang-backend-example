package util

type (
	Response struct {
		Data   interface{} `json:"data,omitempty"`
		Errors []string    `json:"errors,omitempty"`
	}
)
