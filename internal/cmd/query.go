package cmd

import (
	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
)

type QueryCmd struct {
	Query string `arg:"" help:"SQL-like query string (e.g. SELECT * FROM Invoice WHERE Balance > '0')."`
}

func (c *QueryCmd) Run(g *Globals) error {
	if c.Query == "" {
		return errfmt.Usage("query string required")
	}

	if g.CLI.DryRun {
		output.Hint("[dry-run] GET /v3/company/{id}/query?query=%s", c.Query)
		return nil
	}

	client, _, err := g.NewAPIClient()
	if err != nil {
		return err
	}
	result, err := client.Query(g.Ctx, c.Query)
	if err != nil {
		return err
	}
	return WriteOutput(g.Ctx, result)
}
