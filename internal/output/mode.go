package output

import (
	"context"
	"os"
)

type Mode int

const (
	ModeHuman Mode = iota
	ModeJSON
	ModePlain
)

type ctxKey struct{}

type Options struct {
	Mode        Mode
	ResultsOnly bool
	Select      []string
	Pretty      bool
}

func NewOptions(jsonFlag, plainFlag, resultsOnly bool, selectFields []string) Options {
	mode := ModeHuman
	switch {
	case jsonFlag:
		mode = ModeJSON
	case plainFlag:
		mode = ModePlain
	default:
		if os.Getenv("QBO_AUTO_JSON") == "1" && !isTerminal() {
			mode = ModeJSON
		}
	}
	return Options{
		Mode:        mode,
		ResultsOnly: resultsOnly,
		Select:      selectFields,
		Pretty:      isTerminal(),
	}
}

func WithOptions(ctx context.Context, opts Options) context.Context {
	return context.WithValue(ctx, ctxKey{}, opts)
}

func GetOptions(ctx context.Context) Options {
	if v, ok := ctx.Value(ctxKey{}).(Options); ok {
		return v
	}
	return Options{Mode: ModeHuman, Pretty: true}
}

func isTerminal() bool {
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}
