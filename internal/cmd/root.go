package cmd

import (
	"context"
	"os"
	"sync"

	"github.com/voska/qbo-cli/internal/api"
	"github.com/voska/qbo-cli/internal/auth"
	"github.com/voska/qbo-cli/internal/config"
	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
	"golang.org/x/oauth2"
)

type CLI struct {
	JSON        bool         `short:"j" name:"json" help:"Output JSON to stdout." aliases:"machine"`
	Plain       bool         `short:"p" name:"plain" help:"Output TSV, no color." aliases:"tsv"`
	Select      []string     `name:"select" help:"Project output to fields (dot paths)." aliases:"fields"`
	ResultsOnly bool         `name:"results-only" help:"Strip pagination metadata."`
	CompanyID   string       `name:"company-id" env:"QBO_COMPANY_ID" help:"Target company realm ID."`
	Sandbox     bool         `help:"Use sandbox environment."`
	DryRun      bool         `short:"n" name:"dry-run" help:"Show what would happen." aliases:"noop,preview"`
	NoInput     bool         `name:"no-input" help:"Never prompt; fail if confirmation needed."`
	Force       bool         `help:"Skip confirmations." aliases:"yes,assume-yes"`
	Verbose     bool         `short:"v" help:"Verbose output to stderr."`
	MinorVer    int          `name:"minor-version" default:"75" help:"QBO API minor version."`
	Auth        AuthCmd      `cmd:"" help:"Authentication (login, logout, status, refresh)."`
	Query       QueryCmd     `cmd:"" help:"Run a raw query against the QBO API."`
	Get         GetCmd       `cmd:"" help:"Read a single entity by ID."`
	Create      CreateCmd    `cmd:"" help:"Create an entity from JSON."`
	Update      UpdateCmd    `cmd:"" help:"Update an entity."`
	Delete      DeleteCmd    `cmd:"" help:"Delete an entity by ID."`
	List        ListCmd      `cmd:"" help:"List entities (sugar for query)."`
	Batch       BatchCmd     `cmd:"" help:"Run batch operations."`
	CDC         CDCCmd       `cmd:"" name:"cdc" help:"Change data capture polling."`
	Report      ReportCmd    `cmd:"" help:"Run a QBO report."`
	Company     CompanyCmd   `cmd:"" help:"Company info and switching."`
	Attach      AttachCmd    `cmd:"" help:"Upload a file attachment to an entity."`
	Download    DownloadCmd  `cmd:"" help:"Download an attachment file."`
	Schema      SchemaCmd    `cmd:"" help:"Dump CLI schema as JSON for agent introspection."`
	ExitCodes   ExitCodesCmd `cmd:"" name:"exit-codes" help:"Print exit code table."`
}

type Globals struct {
	Ctx     context.Context
	Config  *config.Config
	OutOpts output.Options
	CLI     *CLI
	Version string

	ccOnce sync.Once
	cc     auth.ClientCreds
}

// keyringCreds lazily loads the app-level OAuth client credentials from the
// keyring, once per process. It is only consulted when the corresponding env
// var is unset (see the resolvers), so env-configured callers never open the
// keyring. A read failure degrades to empty creds — warned once — so the
// env/config resolution tiers still apply.
func (g *Globals) keyringCreds() auth.ClientCreds {
	g.ccOnce.Do(func() {
		c, ok, err := auth.LoadClientCreds()
		switch {
		case err != nil:
			output.Warn("could not read client credentials from keyring: %v", err)
		case ok:
			g.cc = c
		}
	})
	return g.cc
}

// ClientID/ClientSecret/RedirectURI resolve credentials across all tiers:
// env var → keyring → config file. The keyring closures are lazy: they only
// run when the env var is unset, so an env-only setup never touches the keyring.
func (g *Globals) ClientID() string {
	return g.Config.ResolveClientID(func() string { return g.keyringCreds().ClientID })
}

func (g *Globals) ClientSecret() string {
	return g.Config.ResolveClientSecret(func() string { return g.keyringCreds().ClientSecret })
}

func (g *Globals) RedirectURI() string {
	return g.Config.ResolveRedirectURI(func() string { return g.keyringCreds().RedirectURI })
}

func NewGlobals(ctx context.Context, cli *CLI) (*Globals, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	opts := output.NewOptions(cli.JSON, cli.Plain, cli.ResultsOnly, cli.Select)
	return &Globals{
		Ctx:     output.WithOptions(ctx, opts),
		Config:  cfg,
		OutOpts: opts,
		CLI:     cli,
	}, nil
}

func (g *Globals) ResolveCompanyID() (string, error) {
	id := g.Config.ResolveCompanyID(g.CLI.CompanyID)
	if id == "" {
		return "", errfmt.Config("no company ID — set --company-id, QBO_COMPANY_ID, or run: qbo auth login")
	}
	return id, nil
}

func (g *Globals) NewAPIClient() (*api.Client, string, error) {
	realmID, err := g.ResolveCompanyID()
	if err != nil {
		return nil, "", err
	}
	token, err := g.loadAndRefreshToken(realmID)
	if err != nil {
		return nil, "", err
	}
	return api.NewClient(token, realmID, g.CLI.Sandbox, g.CLI.MinorVer), realmID, nil
}

func (g *Globals) loadAndRefreshToken(realmID string) (*oauth2.Token, error) {
	token, err := auth.LoadToken(realmID)
	if err != nil {
		return nil, err
	}
	if auth.IsTokenExpired(token) {
		clientID := g.ClientID()
		clientSecret := g.ClientSecret()
		if clientID == "" || clientSecret == "" {
			return nil, errfmt.Config("client credentials required for token refresh — run: qbo auth set-client (or set QBO_CLIENT_ID and QBO_CLIENT_SECRET)")
		}
		newToken, err := auth.RefreshAccessToken(g.Ctx, clientID, clientSecret, token)
		if err != nil {
			return nil, err
		}
		if err := auth.StoreToken(realmID, newToken); err != nil {
			output.Warn("could not persist refreshed token: %v", err)
		}
		return newToken, nil
	}
	return token, nil
}

func WriteOutput(ctx context.Context, data any) error {
	return output.Write(ctx, data)
}

func DryRunLog(cli *CLI, method, endpoint string, body []byte) {
	output.Hint("[dry-run] %s %s", method, endpoint)
	if body != nil {
		_, _ = os.Stderr.Write(body)
		_, _ = os.Stderr.Write([]byte("\n"))
	}
}
