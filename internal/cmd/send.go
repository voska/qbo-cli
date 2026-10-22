package cmd

import (
	"fmt"
	"net/mail"
	"net/url"
	"os"

	"github.com/voska/qbo-cli/internal/api"
	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
	"golang.org/x/term"
)

type SendCmd struct {
	Entity string  `arg:"" help:"Entity type (invoice only)."`
	ID     string  `arg:"" help:"Numeric invoice Id, not its document number."`
	To     *string `name:"to" placeholder:"EMAIL" help:"Recipient email; replaces the invoice's BillEmail.Address."`
}

func (c *SendCmd) Run(g *Globals) error {
	entity, ok := api.LookupEntity(c.Entity)
	if !ok || entity.Endpoint != "invoice" {
		return errfmt.Usage("send supports invoice only")
	}
	if !numericID.MatchString(c.ID) {
		return errfmt.Usage("id must be the numeric invoice Id, got: " + c.ID)
	}
	to := ""
	recipient := "the invoice's BillEmail"
	if c.To != nil {
		address, err := mail.ParseAddress(*c.To)
		if err != nil || address.Address != *c.To {
			return errfmt.Usage("--to must be a single email address without a display name")
		}
		to = *c.To
		recipient = to + " (replaces the invoice's BillEmail.Address)"
	}

	recipient += " (plus any invoice Cc/Bcc)"

	if g.CLI.DryRun {
		endpoint := "/v3/company/{id}/invoice/" + c.ID + "/send"
		if to != "" {
			endpoint += "?sendTo=" + url.QueryEscape(to)
		}
		DryRunLog(g.CLI, "POST", endpoint, nil)
		output.Hint("Emails a third party: %s. Sends are never retried.", recipient)
		return nil
	}

	if !g.CLI.Force {
		if g.CLI.NoInput || !term.IsTerminal(int(os.Stdin.Fd())) {
			return errfmt.Usage("refusing to email a third party without --force in noninteractive mode")
		}
		if !confirmSend(c.ID, recipient) {
			output.Hint("cancelled; no send submitted")
			return nil
		}
	}

	client, _, err := g.NewAPIClient()
	if err != nil {
		return err
	}
	result, err := client.SendInvoice(g.Ctx, c.ID, to)
	if err != nil {
		return err
	}
	if opts := output.GetOptions(g.Ctx); opts.Mode != output.ModeJSON && len(opts.Select) == 0 {
		fields := []string{"Id", "DocNumber", "EmailStatus", "BillEmail.Address", "DeliveryInfo.DeliveryTime"}
		summary := output.ProjectFields(result["Invoice"], fields).(map[string]any)
		for _, field := range fields {
			if summary[field] == nil {
				summary[field] = ""
			}
		}
		return WriteOutput(g.Ctx, summary)
	}
	return WriteOutput(g.Ctx, result)
}

func confirmSend(id, recipient string) bool {
	fmt.Fprintf(os.Stderr, "Email invoice %s to %s? This emails a third party. [y/N] ", id, recipient)
	var confirm string
	_, err := fmt.Fscanln(os.Stdin, &confirm)
	return err == nil && (confirm == "y" || confirm == "Y")
}
