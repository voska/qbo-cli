package output

import (
	"encoding/json"
	"io"
	"os"
)

func WriteJSON(data any, pretty bool) error {
	return writeJSON(os.Stdout, data, pretty)
}

func writeJSON(w io.Writer, data any, pretty bool) error {
	enc := json.NewEncoder(w)
	if pretty {
		enc.SetIndent("", "  ")
	}
	enc.SetEscapeHTML(false)
	return enc.Encode(data)
}
