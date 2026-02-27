package cmd

import (
	"github.com/voska/qbo-cli/internal/output"
)

type ReportCmd struct {
	Type      string `arg:"" help:"Report type (e.g. ProfitAndLoss, BalanceSheet, CashFlow, GeneralLedger)." enum:"ProfitAndLoss,BalanceSheet,CashFlow,TrialBalance,GeneralLedger,AccountList,profit-and-loss,balance-sheet,cash-flow,trial-balance,general-ledger,account-list"`
	StartDate string `name:"start-date" help:"Report start date (YYYY-MM-DD)."`
	EndDate   string `name:"end-date" help:"Report end date (YYYY-MM-DD)."`
}

var reportAliases = map[string]string{
	"profit-and-loss": "ProfitAndLoss",
	"balance-sheet":   "BalanceSheet",
	"cash-flow":       "CashFlow",
	"trial-balance":   "TrialBalance",
	"general-ledger":  "GeneralLedger",
	"account-list":    "AccountList",
}

func (c *ReportCmd) Run(g *Globals) error {
	reportType := c.Type
	if alias, ok := reportAliases[reportType]; ok {
		reportType = alias
	}

	params := map[string]string{}
	if c.StartDate != "" {
		params["start_date"] = c.StartDate
	}
	if c.EndDate != "" {
		params["end_date"] = c.EndDate
	}

	if g.CLI.DryRun {
		output.Hint("[dry-run] GET /v3/company/{id}/reports/%s", reportType)
		return nil
	}

	client, _, err := g.NewAPIClient()
	if err != nil {
		return err
	}
	result, err := client.Report(g.Ctx, reportType, params)
	if err != nil {
		return err
	}
	return WriteOutput(g.Ctx, result)
}
