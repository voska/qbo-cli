package cmd

import (
	"fmt"
	"os"

	"github.com/voska/qbo-cli/internal/output"
)

type InvoiceCmd struct {
	PDF InvoicePDFCmd `cmd:"" name:"pdf" help:"Download a QBO-rendered invoice PDF."`
}

type InvoicePDFCmd struct {
	ID  string `arg:"" help:"Invoice ID."`
	Out string `short:"o" name:"out" required:"" help:"Output PDF file path."`
}

func (c *InvoicePDFCmd) Run(g *Globals) error {
	if g.CLI.DryRun {
		output.Hint("[dry-run] GET /v3/company/{id}/invoice/%s/pdf", c.ID)
		output.Hint("[dry-run] write PDF to %s", c.Out)
		return nil
	}

	client, _, err := g.NewAPIClient()
	if err != nil {
		return err
	}
	pdf, err := client.InvoicePDF(g.Ctx, c.ID)
	if err != nil {
		return err
	}
	if err := os.WriteFile(c.Out, pdf, 0o644); err != nil {
		return fmt.Errorf("cannot write PDF: %w", err)
	}

	switch g.OutOpts.Mode {
	case output.ModeJSON:
		return WriteOutput(g.Ctx, map[string]any{
			"invoice_id": c.ID,
			"out":        c.Out,
			"bytes":      len(pdf),
		})
	case output.ModePlain:
		fmt.Fprintf(os.Stdout, "%s\t%s\t%d\n", c.ID, c.Out, len(pdf))
	default:
		output.Success("wrote invoice PDF to %s", c.Out)
	}
	return nil
}
