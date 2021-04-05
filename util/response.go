package util

type (
	Response struct {
		Data   interface{} `json:"data"`
		Errors []string    `json:"errors,omitempty"`
	}
)
