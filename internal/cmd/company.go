package cmd

import (
	"github.com/voska/qbo-cli/internal/api"
	"github.com/voska/qbo-cli/internal/output"
)

type CompanyCmd struct {
	Info   CompanyInfoCmd   `cmd:"" help:"Show company info from QBO API."`
	List   CompanyListCmd   `cmd:"" help:"List configured companies."`
	Switch CompanySwitchCmd `cmd:"" help:"Set default company."`
}

type CompanyInfoCmd struct{}

func (c *CompanyInfoCmd) Run(g *Globals) error {
	if g.CLI.DryRun {
		output.Hint("[dry-run] GET /v3/company/{id}/companyinfo/{id}")
		return nil
	}

	client, realmID, err := g.NewAPIClient()
	if err != nil {
		return err
	}

	q := api.NewQuery("CompanyInfo").MaxResults(1)
	result, err := client.Query(g.Ctx, q.Build())
	if err != nil {
		return err
	}

	if g.CLI.Verbose {
		output.Hint("company: %s", realmID)
	}
	return WriteOutput(g.Ctx, result)
}

type CompanyListCmd struct{}

func (c *CompanyListCmd) Run(g *Globals) error {
	if len(g.Config.Companies) == 0 {
		output.Hint("no companies configured — run: qbo auth login")
		return nil
	}

	companies := make([]map[string]any, len(g.Config.Companies))
	for i, co := range g.Config.Companies {
		companies[i] = map[string]any{
			"realm_id":    co.RealmID,
			"name":        co.CompanyName,
			"environment": co.Environment,
			"default":     co.RealmID == g.Config.DefaultCompany,
		}
	}
	return WriteOutput(g.Ctx, companies)
}

type CompanySwitchCmd struct {
	RealmID string `arg:"" help:"Realm ID to set as default."`
}

func (c *CompanySwitchCmd) Run(g *Globals) error {
	g.Config.DefaultCompany = c.RealmID
	if err := g.Config.Save(); err != nil {
		return err
	}
	output.Success("default company set to %s", c.RealmID)
	return nil
}
