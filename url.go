package wichess

import (
	"fmt"
	"net/url"
	"strconv"
)

func parseURLQuery(from url.Values, queryKey string) (string, error) {
	q, has := from[queryKey]
	if has == false {
		return "", fmt.Errorf("%v not in URL queries", queryKey)
	}

	if len(q) != 1 {
		return "", fmt.Errorf("%v has %v values instead of 1", queryKey, len(q))
	}

	return q[0], nil
}
