package main

import (
	"errors"
	"os"

	"github.com/alecthomas/kong"
	"github.com/voska/qbo-cli/internal/cmd"
	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
)

var version = "dev"

func main() {
	var cli cmd.CLI
	parser := kong.Must(&cli,
		kong.Name("qbo"),
		kong.Description("QuickBooks Online CLI for humans and AI agents."),
		kong.UsageOnError(),
	)

	ctx, err := parser.Parse(os.Args[1:])
	if err != nil {
		parser.FatalIfErrorf(err)
	}

	globals, err := cmd.NewGlobals(&cli)
	if err != nil {
		handleError(err)
	}
	globals.Version = version

	err = ctx.Run(globals)
	if err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	var e *errfmt.Error
	if errors.As(err, &e) {
		output.ErrorMsg("%s", e.Error())
		os.Exit(e.Code)
	}
	output.ErrorMsg("%s", err.Error())
	os.Exit(errfmt.ExitError)
}
