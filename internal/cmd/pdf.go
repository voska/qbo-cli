package cmd

import (
	"os"
	"regexp"

	"github.com/voska/qbo-cli/internal/api"
	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
)

type PDFCmd struct {
	Entity string `arg:"" help:"Transaction type: invoice, estimate, salesreceipt, creditmemo or purchaseorder."`
	ID     string `arg:"" help:"Transaction ID (the Id, not the document number)."`
	Output string `name:"output" short:"o" help:"Save to this file path. Defaults to <entity>-<id>.pdf."`
}

var numericID = regexp.MustCompile(`^[0-9]+$`)

func (c *PDFCmd) Run(g *Globals) error {
	entity, ok := api.LookupEntity(c.Entity)
	if !ok {
		return errfmt.Usage("unknown entity: " + c.Entity)
	}
	if !api.SupportsPDF(entity.Endpoint) {
		return errfmt.Usage(entity.Name + " has no PDF rendering — use invoice, estimate, salesreceipt, creditmemo or purchaseorder")
	}
	// The id is interpolated into the URL path, and a document number passed by
	// mistake would fetch a different transaction that happens to have that Id.
	if !numericID.MatchString(c.ID) {
		return errfmt.Usage("id must be the numeric transaction Id, got: " + c.ID)
	}

	savePath := c.Output
	if savePath == "" {
		savePath = entity.Endpoint + "-" + c.ID + ".pdf"
	}

	if g.CLI.DryRun {
		output.Hint("[dry-run] GET /v3/company/{id}/%s/%s/pdf → %s", entity.Endpoint, c.ID, savePath)
		return nil
	}

	client, _, err := g.NewAPIClient()
	if err != nil {
		return err
	}
	pdf, err := client.PDF(g.Ctx, entity.Endpoint, c.ID)
	if err != nil {
		return err
	}
	if err := os.WriteFile(savePath, pdf, 0o600); err != nil {
		return errfmt.Wrap(errfmt.ExitError, "cannot write file", err)
	}

	output.Hint("saved %s (%d bytes)", savePath, len(pdf))
	return WriteOutput(g.Ctx, map[string]any{"entity": entity.Name, "id": c.ID, "path": savePath, "bytes": len(pdf)})
}
