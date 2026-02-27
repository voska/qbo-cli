package output

import (
	"strings"
)

func ProjectFields(data any, fields []string) any {
	if len(fields) == 0 {
		return data
	}
	switch v := data.(type) {
	case []any:
		result := make([]any, len(v))
		for i, item := range v {
			result[i] = projectOne(item, fields)
		}
		return result
	case map[string]any:
		return projectOne(v, fields)
	default:
		return data
	}
}

func projectOne(data any, fields []string) any {
	m, ok := data.(map[string]any)
	if !ok {
		return data
	}
	result := make(map[string]any, len(fields))
	for _, field := range fields {
		if val := getPath(m, field); val != nil {
			result[field] = val
		}
	}
	return result
}

func getPath(m map[string]any, path string) any {
	parts := strings.Split(path, ".")
	var current any = m
	for _, part := range parts {
		cm, ok := current.(map[string]any)
		if !ok {
			return nil
		}
		current, ok = cm[part]
		if !ok {
			return nil
		}
	}
	return current
}

func StripMetadata(data any) any {
	m, ok := data.(map[string]any)
	if !ok {
		return data
	}
	if qr, ok := m["QueryResponse"]; ok {
		if qrMap, ok := qr.(map[string]any); ok {
			for k, v := range qrMap {
				if k != "startPosition" && k != "maxResults" && k != "totalCount" {
					if arr, ok := v.([]any); ok {
						return arr
					}
				}
			}
		}
	}
	return data
}
