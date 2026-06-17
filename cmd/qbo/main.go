package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

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

	kctx, err := parser.Parse(os.Args[1:])
	if err != nil {
		parser.FatalIfErrorf(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	globals, err := cmd.NewGlobals(ctx, &cli)
	if err != nil {
		handleError(err)
	}
	globals.Version = version

	if err := kctx.Run(globals); err != nil {
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
