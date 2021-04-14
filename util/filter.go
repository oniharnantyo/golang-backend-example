package util

type (
	Filter struct {
		Limit  int    `json:"limit" form:"limit"`
		Offset int    `json:"offset" form:"offset"`
		Search string `json:"search" form:"search"`
		Order  string `json:"order" form:"order"`
	}
)
