package cmd

import (
	"github.com/voska/qbo-cli/internal/api"
	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
)

type GetCmd struct {
	Entity string `arg:"" help:"Entity type (e.g. invoice, customer)."`
	ID     string `arg:"" help:"Entity ID."`
}

func (c *GetCmd) Run(g *Globals) error {
	entity, ok := api.LookupEntity(c.Entity)
	if !ok {
		return errfmt.Usage("unknown entity: " + c.Entity)
	}

	if g.CLI.DryRun {
		output.Hint("[dry-run] GET /v3/company/{id}/%s/%s", entity.Endpoint, c.ID)
		return nil
	}

	client, _, err := g.NewAPIClient()
	if err != nil {
		return err
	}
	result, err := client.Read(g.Ctx, entity.Endpoint, c.ID)
	if err != nil {
		return err
	}
	return WriteOutput(g.Ctx, result)
}
