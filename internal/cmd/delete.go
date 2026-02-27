package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/voska/qbo-cli/internal/api"
	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
)

type DeleteCmd struct {
	Entity string `arg:"" help:"Entity type (e.g. invoice, customer)."`
	ID     string `arg:"" help:"Entity ID."`
}

func (c *DeleteCmd) Run(g *Globals) error {
	entity, ok := api.LookupEntity(c.Entity)
	if !ok {
		return errfmt.Usage("unknown entity: " + c.Entity)
	}
	if !entity.Deletable {
		return errfmt.Usage(entity.Name + " cannot be deleted")
	}

	if g.CLI.DryRun {
		output.Hint("[dry-run] POST /v3/company/{id}/%s?operation=delete (id=%s)", entity.Endpoint, c.ID)
		return nil
	}

	if !g.CLI.Force && !g.CLI.NoInput {
		fmt.Fprintf(os.Stderr, "Delete %s %s? [y/N] ", entity.Name, c.ID)
		var confirm string
		fmt.Scanln(&confirm)
		if confirm != "y" && confirm != "Y" {
			output.Hint("cancelled")
			return nil
		}
	}

	client, _, err := g.NewAPIClient()
	if err != nil {
		return err
	}

	body, _ := json.Marshal(map[string]any{
		"Id":       c.ID,
		"SyncToken": "0",
	})

	result, err := client.Delete(g.Ctx, entity.Endpoint, body)
	if err != nil {
		return err
	}
	return WriteOutput(g.Ctx, result)
}
