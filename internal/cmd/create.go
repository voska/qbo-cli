package cmd

import (
	"io"
	"os"

	"github.com/voska/qbo-cli/internal/api"
	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
)

type CreateCmd struct {
	Entity string `arg:"" help:"Entity type (e.g. invoice, customer)."`
	File   string `name:"file" short:"f" help:"JSON file path, or - for stdin." type:"existingfile"`
}

func (c *CreateCmd) Run(g *Globals) error {
	entity, ok := api.LookupEntity(c.Entity)
	if !ok {
		return errfmt.Usage("unknown entity: " + c.Entity)
	}
	if !entity.Creatable {
		return errfmt.Usage(entity.Name + " cannot be created")
	}

	body, err := readInput(c.File)
	if err != nil {
		return err
	}

	if g.CLI.DryRun {
		output.Hint("[dry-run] POST /v3/company/{id}/%s", entity.Endpoint)
		DryRunLog(g.CLI, "POST", entity.Endpoint, body)
		return nil
	}

	client, _, err := g.NewAPIClient()
	if err != nil {
		return err
	}
	result, err := client.Create(g.Ctx, entity.Endpoint, body)
	if err != nil {
		return err
	}
	return WriteOutput(g.Ctx, result)
}

func readInput(file string) ([]byte, error) {
	if file == "-" || file == "" {
		fi, _ := os.Stdin.Stat()
		if (fi.Mode() & os.ModeCharDevice) == 0 {
			return io.ReadAll(os.Stdin)
		}
		return nil, errfmt.Usage("provide --file or pipe JSON via stdin")
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot read file", err)
	}
	return data, nil
}
