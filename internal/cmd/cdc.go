package cmd

import (
	"strings"

	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
)

type CDCCmd struct {
	Entities string `name:"entities" required:"" help:"Comma-separated entity names (e.g. Customer,Invoice)."`
	Since    string `name:"since" required:"" help:"Changed since timestamp (ISO 8601, e.g. 2025-02-20T00:00:00Z)."`
}

func (c *CDCCmd) Run(g *Globals) error {
	names := strings.Split(c.Entities, ",")
	if len(names) == 0 {
		return errfmt.Usage("at least one entity required")
	}

	if g.CLI.DryRun {
		output.Hint("[dry-run] GET /v3/company/{id}/cdc?entities=%s&changedSince=%s", c.Entities, c.Since)
		return nil
	}

	client, _, err := g.NewAPIClient()
	if err != nil {
		return err
	}
	result, err := client.CDC(g.Ctx, names, c.Since)
	if err != nil {
		return err
	}
	return WriteOutput(g.Ctx, result)
}
