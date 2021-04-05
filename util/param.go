package util

import (
	"html"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/gorilla/schema"
)

func ParseQueryParams(r *http.Request) (Filter, error) {
	var filter Filter
	err := schema.NewDecoder().Decode(&filter, r.Form)
	if err != nil {
		return Filter{}, err
	}

	if strings.ToLower(filter.Order) != "asc" && strings.ToLower(filter.Order) != "desc" {
		return Filter{}, errors.Wrap(err, "Invalid order value")
	}

	// Escaping value to prevent SQL Injection
	filter.Search = html.EscapeString(filter.Search)

	return filter, nil
}
