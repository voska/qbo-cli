package cmd

import (
	"github.com/voska/qbo-cli/internal/api"
	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
)

type UpdateCmd struct {
	Entity string `arg:"" help:"Entity type (e.g. invoice, customer)."`
	ID     string `arg:"" help:"Entity ID."`
	File   string `name:"file" short:"f" help:"JSON file path, or - for stdin." type:"existingfile"`
	Sparse bool   `help:"Perform a sparse (partial) update."`
}

func (c *UpdateCmd) Run(g *Globals) error {
	entity, ok := api.LookupEntity(c.Entity)
	if !ok {
		return errfmt.Usage("unknown entity: " + c.Entity)
	}
	if !entity.Updatable {
		return errfmt.Usage(entity.Name + " cannot be updated")
	}

	body, err := readInput(c.File)
	if err != nil {
		return err
	}

	if g.CLI.DryRun {
		op := "full"
		if c.Sparse {
			op = "sparse"
		}
		output.Hint("[dry-run] POST /v3/company/{id}/%s (%s update, id=%s)", entity.Endpoint, op, c.ID)
		DryRunLog(g.CLI, "POST", entity.Endpoint, body)
		return nil
	}

	client, _, err := g.NewAPIClient()
	if err != nil {
		return err
	}
	result, err := client.Update(g.Ctx, entity.Endpoint, body, c.Sparse)
	if err != nil {
		return err
	}
	return WriteOutput(g.Ctx, result)
}
