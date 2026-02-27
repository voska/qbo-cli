package output

import (
	"fmt"
	"os"
	"strings"
)

func WritePlain(headers []string, rows [][]string) {
	if len(headers) > 0 {
		fmt.Fprintln(os.Stdout, strings.Join(headers, "\t"))
	}
	for _, row := range rows {
		fmt.Fprintln(os.Stdout, strings.Join(row, "\t"))
	}
}
