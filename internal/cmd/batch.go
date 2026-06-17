package cmd

import (
	"encoding/json"

	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
)

type BatchCmd struct {
	File string `name:"file" short:"f" required:"" help:"JSON file with batch operations, or - for stdin." type:"existingfile"`
}

func (c *BatchCmd) Run(g *Globals) error {
	body, err := readInput(c.File)
	if err != nil {
		return err
	}

	var items []json.RawMessage
	if err := json.Unmarshal(body, &items); err != nil {
		return errfmt.Wrap(errfmt.ExitUsage, "invalid batch JSON (expected a JSON array of batch items)", err)
	}
	if len(items) == 0 {
		return errfmt.Usage("batch file contains no items")
	}

	if g.CLI.DryRun {
		output.Hint("[dry-run] POST /v3/company/{id}/batch (%d items)", len(items))
		DryRunLog(g.CLI, "POST", "batch", body)
		return nil
	}

	client, _, err := g.NewAPIClient()
	if err != nil {
		return err
	}

	result, err := client.Batch(g.Ctx, items)
	if err != nil {
		return err
	}
	return WriteOutput(g.Ctx, result)
}
