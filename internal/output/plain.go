package output

import (
	"fmt"
	"os"
	"strings"
)

func WritePlain(headers []string, rows [][]string) {
	if len(headers) > 0 {
		_, _ = fmt.Fprintln(os.Stdout, strings.Join(headers, "\t"))
	}
	for _, row := range rows {
		_, _ = fmt.Fprintln(os.Stdout, strings.Join(row, "\t"))
	}
}
