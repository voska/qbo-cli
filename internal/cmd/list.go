package cmd

import (
	"github.com/voska/qbo-cli/internal/api"
	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
)

type ListCmd struct {
	Entity  string `arg:"" help:"Entity type (e.g. invoices, customers)."`
	Where   string `name:"where" short:"w" help:"WHERE clause for filtering."`
	OrderBy string `name:"order-by" help:"ORDERBY clause."`
	Limit   int    `name:"limit" short:"l" default:"100" help:"Max results to return."`
	Offset  int    `name:"offset" default:"1" help:"Start position."`
}

func (c *ListCmd) Run(g *Globals) error {
	entity, ok := api.LookupEntity(c.Entity)
	if !ok {
		return errfmt.Usage("unknown entity: " + c.Entity)
	}
	if !entity.Queryable {
		return errfmt.Usage(entity.Name + " cannot be queried")
	}

	q := api.NewQuery(entity.Name).
		Where(c.Where).
		OrderBy(c.OrderBy).
		MaxResults(c.Limit).
		StartPosition(c.Offset)

	query := q.Build()

	if g.CLI.DryRun {
		output.Hint("[dry-run] query: %s", query)
		return nil
	}

	client, _, err := g.NewAPIClient()
	if err != nil {
		return err
	}
	result, err := client.Query(g.Ctx, query)
	if err != nil {
		return err
	}
	return WriteOutput(g.Ctx, result)
}
