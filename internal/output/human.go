package output

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/muesli/termenv"
)

var profile = termenv.ColorProfile()

func WriteTable(headers []string, rows [][]string) {
	writeTable(os.Stdout, headers, rows)
}

func writeTable(w io.Writer, headers []string, rows [][]string) {
	if len(headers) == 0 && len(rows) == 0 {
		return
	}

	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = len(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if i < len(widths) && len(cell) > widths[i] {
				widths[i] = len(cell)
			}
		}
	}

	bold := termenv.Style{}.Bold()
	if len(headers) > 0 {
		parts := make([]string, len(headers))
		for i, h := range headers {
			parts[i] = bold.Styled(padRight(h, widths[i]))
		}
		fmt.Fprintln(w, strings.Join(parts, "  "))
		sep := make([]string, len(headers))
		for i, width := range widths {
			sep[i] = strings.Repeat("─", width)
		}
		fmt.Fprintln(w, strings.Join(sep, "  "))
	}

	for _, row := range rows {
		parts := make([]string, len(headers))
		for i := range headers {
			cell := ""
			if i < len(row) {
				cell = row[i]
			}
			parts[i] = padRight(cell, widths[i])
		}
		fmt.Fprintln(w, strings.Join(parts, "  "))
	}
}

func padRight(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}

func Hint(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	styled := termenv.String(msg).Foreground(profile.Color("8"))
	fmt.Fprintln(os.Stderr, styled)
}

func Success(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	styled := termenv.String("✓ " + msg).Foreground(profile.Color("2"))
	fmt.Fprintln(os.Stderr, styled)
}

func Warn(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	styled := termenv.String("⚠ " + msg).Foreground(profile.Color("3"))
	fmt.Fprintln(os.Stderr, styled)
}

func ErrorMsg(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	styled := termenv.String("✗ " + msg).Foreground(profile.Color("1"))
	fmt.Fprintln(os.Stderr, styled)
}
