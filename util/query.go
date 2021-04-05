package util

import "fmt"

func BuildFilterQuery(filters []string) string {
	var filtersQuery string
	if len(filters) != 0 {
		filtersQuery = "WHERE"
		if len(filters) == 1 {
			filtersQuery = fmt.Sprintf(`%s %s`, filtersQuery, filters[0])
		} else {
			for i, filter := range filters {
				if i == 0 {
					filtersQuery = fmt.Sprintf(`%s %s`, filtersQuery, filter)
				}
				if i > 0 {
					filtersQuery = fmt.Sprintf(`%s AND %s`, filtersQuery, filter)
				}
			}
		}
	}
	return filtersQuery
}
