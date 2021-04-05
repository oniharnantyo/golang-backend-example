package util

type (
	Filter struct {
		Limit  int    `json:"limit" schema:"limit"`
		Offset int    `json:"offset" schema:"offset"`
		Search string `json:"search" schema:"search"`
		Order  string `json:"order" schema:"order"`
	}
)
