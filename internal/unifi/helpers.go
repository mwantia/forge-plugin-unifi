package unifi

import (
	"fmt"
	"net/http"
)

func LocalBaseUrl(address string) string {
	return address + "/proxy/network/integration/v1"
}

func SetHttpAPIHeaders(request *http.Request, apiKey string) {
	request.Header.Set("Accept", "application/json")
	request.Header.Set("X-API-Key", apiKey)
}

func PaginatedResult[T any](key string, data []T, r *apiResponse[T]) map[string]any {
	return map[string]any{
		key:           data,
		"count":       r.Count,
		"total_count": r.TotalCount,
		"offset":      r.Offset,
		"limit":       r.Limit,
	}
}

func PaginationQuery(args map[string]any) string {
	limit, hasLimit := IntArg(args, "limit")
	offset, hasOffset := IntArg(args, "offset")
	if !hasLimit && !hasOffset {
		return ""
	}
	q := "?"
	if hasLimit {
		q += fmt.Sprintf("limit=%d", limit)
	}
	if hasOffset {
		if hasLimit {
			q += "&"
		}
		q += fmt.Sprintf("offset=%d", offset)
	}
	return q
}

func IntArg(args map[string]any, key string) (int, bool) {
	v, ok := args[key]
	if !ok || v == nil {
		return 0, false
	}
	switch n := v.(type) {
	case float64:
		return int(n), true
	case int:
		return n, true
	}
	return 0, false
}
