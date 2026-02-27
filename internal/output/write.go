package output

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
)

func Write(ctx context.Context, data any) error {
	opts := GetOptions(ctx)

	if opts.ResultsOnly {
		data = StripMetadata(data)
	}
	if len(opts.Select) > 0 {
		data = ProjectFields(data, opts.Select)
	}

	switch opts.Mode {
	case ModeJSON:
		return WriteJSON(data, opts.Pretty)
	case ModePlain:
		headers, rows := toTable(data)
		WritePlain(headers, rows)
		return nil
	default:
		headers, rows := toTable(data)
		WriteTable(headers, rows)
		return nil
	}
}

func normalize(data any) any {
	switch data.(type) {
	case map[string]any, []any:
		return data
	default:
		b, err := json.Marshal(data)
		if err != nil {
			return data
		}
		var out any
		if err := json.Unmarshal(b, &out); err != nil {
			return data
		}
		return out
	}
}

func toTable(data any) ([]string, [][]string) {
	data = normalize(data)
	switch v := data.(type) {
	case []any:
		if len(v) == 0 {
			return nil, nil
		}
		headers := collectHeaders(v)
		rows := make([][]string, len(v))
		for i, item := range v {
			m, ok := item.(map[string]any)
			if !ok {
				rows[i] = []string{fmt.Sprint(item)}
				continue
			}
			row := make([]string, len(headers))
			for j, h := range headers {
				if val, ok := m[h]; ok {
					row[j] = fmt.Sprint(val)
				}
			}
			rows[i] = row
		}
		return headers, rows
	case map[string]any:
		headers := sortedKeys(v)
		row := make([]string, len(headers))
		for i, h := range headers {
			row[i] = fmt.Sprint(v[h])
		}
		return headers, [][]string{row}
	default:
		return nil, [][]string{{fmt.Sprint(data)}}
	}
}

func collectHeaders(items []any) []string {
	seen := map[string]bool{}
	var headers []string
	for _, item := range items {
		m, ok := item.(map[string]any)
		if !ok {
			continue
		}
		for k := range m {
			if !seen[k] {
				seen[k] = true
				headers = append(headers, k)
			}
		}
	}
	sort.Strings(headers)
	return headers
}

func sortedKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
