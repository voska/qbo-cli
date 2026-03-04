package cmd

import (
	"fmt"
	"time"

	"github.com/voska/qbo-cli/internal/auth"
	"github.com/voska/qbo-cli/internal/config"
	"github.com/voska/qbo-cli/internal/errfmt"
	"github.com/voska/qbo-cli/internal/output"
)

type AuthCmd struct {
	Login   AuthLoginCmd   `cmd:"" help:"Authenticate with QuickBooks Online."`
	Logout  AuthLogoutCmd  `cmd:"" help:"Remove stored credentials."`
	Status  AuthStatusCmd  `cmd:"" help:"Show current auth status."`
	Refresh AuthRefreshCmd `cmd:"" help:"Force token refresh."`
}

type AuthLoginCmd struct {
	Manual      bool   `help:"Print URL for manual copy instead of opening browser."`
	RedirectURI string `name:"redirect-uri" env:"QBO_REDIRECT_URI" help:"OAuth redirect URI override."`
}

func (c *AuthLoginCmd) Run(g *Globals) error {
	clientID := g.Config.ResolveClientID()
	clientSecret := g.Config.ResolveClientSecret()
	if clientID == "" || clientSecret == "" {
		return errfmt.Config("set QBO_CLIENT_ID and QBO_CLIENT_SECRET before logging in")
	}

	redirectURL := c.RedirectURI
	if redirectURL == "" {
		redirectURL = fmt.Sprintf("http://localhost:%d/callback", auth.DefaultCallbackPort)
	}

	if g.CLI.DryRun {
		url := auth.GetAuthURL(clientID, clientSecret, redirectURL, "STATE")
		output.Hint("[dry-run] would open: %s", url)
		return nil
	}

	result, err := auth.LoginInteractive(g.Ctx, clientID, clientSecret, redirectURL)
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
	return WriteOutput(g.Ctx, status)
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
	clientID := g.Config.ResolveClientID()
	clientSecret := g.Config.ResolveClientSecret()
	if clientID == "" || clientSecret == "" {
		return errfmt.Config("QBO_CLIENT_ID and QBO_CLIENT_SECRET required")
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
