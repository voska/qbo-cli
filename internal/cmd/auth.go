package cmd

import (
	"os"
	"time"

	"github.com/voska/qbo-cli/internal/auth"
	"github.com/voska/qbo-cli/internal/config"
	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
)

type AuthCmd struct {
	Login     AuthLoginCmd     `cmd:"" help:"Authenticate with QuickBooks Online."`
	Logout    AuthLogoutCmd    `cmd:"" help:"Remove stored credentials."`
	Status    AuthStatusCmd    `cmd:"" help:"Show current auth status."`
	Refresh   AuthRefreshCmd   `cmd:"" help:"Force token refresh."`
	SetClient AuthSetClientCmd `cmd:"" name:"set-client" help:"Store OAuth client credentials in the keyring."`
}

type AuthLoginCmd struct {
	Manual      bool   `help:"Print URL for manual copy instead of opening browser."`
	RedirectURI string `name:"redirect-uri" help:"OAuth redirect URI. Required for production (non-localhost). Set via flag, QBO_REDIRECT_URI, or config." env:"QBO_REDIRECT_URI"`
}

func (c *AuthLoginCmd) Run(g *Globals) error {
	clientID := g.ClientID()
	clientSecret := g.ClientSecret()
	if clientID == "" || clientSecret == "" {
		return errfmt.Config("set client credentials first — run: qbo auth set-client (or set QBO_CLIENT_ID and QBO_CLIENT_SECRET)")
	}

	redirectURI := c.RedirectURI
	if redirectURI == "" {
		redirectURI = g.RedirectURI()
	}
	// storedRedirect stays empty when none was configured, so we don't lock the
	// localhost default into the keyring entry.
	storedRedirect := redirectURI
	if redirectURI == "" {
		redirectURI = auth.DefaultRedirectURI()
	}

	if g.CLI.DryRun {
		url := auth.GetAuthURL(clientID, clientSecret, redirectURI, "STATE")
		output.Hint("[dry-run] would open: %s", url)
		return nil
	}

	result, err := auth.LoginInteractive(g.Ctx, clientID, clientSecret, redirectURI)
	if err != nil {
		return err
	}

	if err := auth.StoreToken(result.RealmID, result.Token); err != nil {
		return err
	}

	g.Config.AddOrUpdateCompany(config.Company{
		RealmID:     result.RealmID,
		Environment: resolveEnv(g.CLI),
	})
	if g.Config.DefaultCompany == "" {
		g.Config.DefaultCompany = result.RealmID
	}
	if err := g.Config.Save(); err != nil {
		output.Warn("could not save config: %v", err)
	}

	// Persist the client credentials used so future invocations resolve them
	// from the keyring without QBO_CLIENT_ID/SECRET in the environment.
	if err := auth.StoreClientCreds(auth.ClientCreds{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  storedRedirect,
	}); err != nil {
		output.Warn("authenticated, but could not save client credentials to keyring: %v", err)
	}

	output.Success("authenticated for company %s", result.RealmID)
	return nil
}

type AuthLogoutCmd struct{}

func (c *AuthLogoutCmd) Run(g *Globals) error {
	realmID, err := g.ResolveCompanyID()
	if err != nil {
		return err
	}
	if err := auth.DeleteToken(realmID); err != nil {
		output.Warn("could not remove token: %v", err)
	}
	output.Success("logged out of company %s", realmID)
	return nil
}

type AuthStatusCmd struct{}

func (c *AuthStatusCmd) Run(g *Globals) error {
	realmID, err := g.ResolveCompanyID()
	if err != nil {
		return err
	}
	token, err := auth.LoadToken(realmID)
	if err != nil {
		return err
	}
	status := map[string]any{
		"company_id":    realmID,
		"authenticated": true,
		"token_expiry":  token.Expiry.Format(time.RFC3339),
		"expired":       auth.IsTokenExpired(token),
	}
	co := g.Config.FindCompany(realmID)
	if co != nil {
		status["company_name"] = co.CompanyName
		status["environment"] = co.Environment
	}
	status["client_credentials"] = map[string]any{
		"configured": g.ClientID() != "" && g.ClientSecret() != "",
		"source":     clientCredsSource(g),
	}
	return WriteOutput(g.Ctx, status)
}

// clientCredsSource reports which tier supplied the client credentials, for
// debugging "why can't this agent see QBO" situations.
func clientCredsSource(g *Globals) string {
	switch {
	case os.Getenv("QBO_CLIENT_ID") != "":
		return "env"
	case g.keyringCreds().ClientID != "":
		return "keyring"
	case g.Config.ClientID != "":
		return "config"
	default:
		return "none"
	}
}

type AuthRefreshCmd struct{}

func (c *AuthRefreshCmd) Run(g *Globals) error {
	realmID, err := g.ResolveCompanyID()
	if err != nil {
		return err
	}
	token, err := auth.LoadToken(realmID)
	if err != nil {
		return err
	}
	clientID := g.ClientID()
	clientSecret := g.ClientSecret()
	if clientID == "" || clientSecret == "" {
		return errfmt.Config("client credentials required — run: qbo auth set-client (or set QBO_CLIENT_ID and QBO_CLIENT_SECRET)")
	}
	newToken, err := auth.RefreshAccessToken(g.Ctx, clientID, clientSecret, token)
	if err != nil {
		return err
	}
	if err := auth.StoreToken(realmID, newToken); err != nil {
		return err
	}
	output.Success("token refreshed, expires %s", newToken.Expiry.Format(time.RFC3339))
	return nil
}

func resolveEnv(cli *CLI) string {
	if cli.Sandbox {
		return config.EnvSandbox
	}
	return config.EnvProduction
}

type AuthSetClientCmd struct {
	ClientID     string `name:"client-id" env:"QBO_CLIENT_ID" help:"OAuth client ID (falls back to QBO_CLIENT_ID)."`
	ClientSecret string `name:"client-secret" env:"QBO_CLIENT_SECRET" help:"OAuth client secret (falls back to QBO_CLIENT_SECRET)."`
	RedirectURI  string `name:"redirect-uri" env:"QBO_REDIRECT_URI" help:"Optional OAuth redirect URI."`
	Clear        bool   `help:"Remove stored client credentials from the keyring."`
}

func (c *AuthSetClientCmd) Run(g *Globals) error {
	if c.Clear {
		if g.CLI.DryRun {
			output.Hint("[dry-run] would remove client credentials from keyring")
			return nil
		}
		if err := auth.DeleteClientCreds(); err != nil {
			return err
		}
		output.Success("client credentials removed from keyring")
		return nil
	}

	if c.ClientID == "" || c.ClientSecret == "" {
		return errfmt.Config("provide --client-id and --client-secret (or set QBO_CLIENT_ID and QBO_CLIENT_SECRET)")
	}

	if g.CLI.DryRun {
		output.Hint("[dry-run] would store client credentials in keyring")
		return nil
	}

	if err := auth.StoreClientCreds(auth.ClientCreds{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURI:  c.RedirectURI,
	}); err != nil {
		return err
	}
	output.Success("client credentials stored in keyring")
	return nil
}
