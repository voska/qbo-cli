package cmd

import (
	"fmt"
	"os"

	"github.com/voska/qbo-cli/internal/api"
	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
)

type VoidCmd struct {
	Entity    string `arg:"" help:"Entity type (invoice, payment, salesreceipt, billpayment)."`
	ID        string `arg:"" help:"Entity ID."`
	SyncToken string `name:"sync-token" help:"Expected SyncToken; fail if the current record differs."`
}

func (c *VoidCmd) Run(g *Globals) error {
	entity, ok := api.LookupEntity(c.Entity)
	if !ok {
		return errfmt.Usage("unknown entity: " + c.Entity)
	}
	if entity.VoidStyle == "" {
		return errfmt.Usage(entity.Name + " cannot be voided")
	}

	if g.CLI.DryRun {
		endpoint := "/v3/company/{id}/" + entity.Endpoint + "?" + string(entity.VoidStyle)
		if c.SyncToken == "" {
			output.Hint("[dry-run] GET then POST %s (id=%s; SyncToken read at execution)", endpoint, c.ID)
		} else {
			body, err := entity.VoidBody(c.ID, c.SyncToken)
			if err != nil {
				return err
			}
			plan := fmt.Sprintf("%s (id=%s; require SyncToken=%s)", endpoint, c.ID, c.SyncToken)
			DryRunLog(g.CLI, "GET then POST", plan, body)
		}
		return nil
	}

	if !g.CLI.Force {
		if g.CLI.NoInput {
			return errfmt.Usage("refusing to void without --force in --no-input mode")
		}
		fmt.Fprintf(os.Stderr, "Void %s %s? This can't be undone. [y/N] ", entity.Name, c.ID)
		var confirm string
		_, _ = fmt.Scanln(&confirm)
		if confirm != "y" && confirm != "Y" {
			output.Hint("cancelled")
			return nil
		}
	}

	client, _, err := g.NewAPIClient()
	if err != nil {
		return err
	}
	result, alreadyVoid, err := client.Void(g.Ctx, entity, c.ID, c.SyncToken)
	if err != nil {
		return err
	}
	if alreadyVoid {
		output.Hint("%s %s is already voided; no change", entity.Name, c.ID)
	}
	return WriteOutput(g.Ctx, result)
}
