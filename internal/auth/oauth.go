package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/voska/qbo-cli/internal/errfmt"
	"golang.org/x/oauth2"
)

type AuthResult struct {
	Token   *oauth2.Token
	RealmID string
}

func GenerateState() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func GetAuthURL(clientID, clientSecret, redirectURL, state string) string {
	cfg := OAuthConfig(clientID, clientSecret, redirectURL)
	return cfg.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func ExchangeCode(ctx context.Context, clientID, clientSecret, redirectURL, code string) (*oauth2.Token, error) {
	cfg := OAuthConfig(clientID, clientSecret, redirectURL)
	token, err := cfg.Exchange(ctx, code)
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitAuth, "token exchange failed", err)
	}
	return token, nil
}

func RefreshAccessToken(ctx context.Context, clientID, clientSecret string, token *oauth2.Token) (*oauth2.Token, error) {
	cfg := OAuthConfig(clientID, clientSecret, "")
	src := cfg.TokenSource(ctx, token)
	newToken, err := src.Token()
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitAuth, "token refresh failed", err)
	}
	return newToken, nil
}

const DefaultCallbackPort = 8844

func LoginInteractive(ctx context.Context, clientID, clientSecret, redirectURL string) (*AuthResult, error) {
	addr := fmt.Sprintf("127.0.0.1:%d", DefaultCallbackPort)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, fmt.Sprintf("cannot listen on port %d — is another qbo login running?", DefaultCallbackPort), err)
	}
	if redirectURL == "" {
		redirectURL = fmt.Sprintf("http://localhost:%d/callback", DefaultCallbackPort)
	}

	cfg := OAuthConfig(clientID, clientSecret, redirectURL)
	state := GenerateState()

	resultCh := make(chan *AuthResult, 1)
	errCh := make(chan error, 1)

	mux := http.NewServeMux()
	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			errCh <- errfmt.New(errfmt.ExitAuth, "state mismatch")
			http.Error(w, "state mismatch", http.StatusBadRequest)
			return
		}
		code := r.URL.Query().Get("code")
		realmID := r.URL.Query().Get("realmId")
		if code == "" || realmID == "" {
			errCh <- errfmt.New(errfmt.ExitAuth, "missing code or realmId")
			http.Error(w, "missing parameters", http.StatusBadRequest)
			return
		}
		token, terr := cfg.Exchange(ctx, code)
		if terr != nil {
			errCh <- errfmt.Wrap(errfmt.ExitAuth, "token exchange failed", terr)
			http.Error(w, "exchange failed", http.StatusInternalServerError)
			return
		}
		_, _ = fmt.Fprint(w, "<html><body><h2>Authenticated!</h2><p>You can close this window.</p></body></html>")
		resultCh <- &AuthResult{Token: token, RealmID: realmID}
	})

	server := &http.Server{Handler: mux}
	go func() {
		if serr := server.Serve(listener); serr != http.ErrServerClosed {
			errCh <- serr
		}
	}()
	defer func() { _ = server.Close() }()

	authURL := cfg.AuthCodeURL(state, oauth2.AccessTypeOffline)
	fmt.Fprintf(os.Stderr, "Open this URL in your browser:\n\n  %s\n\n", authURL)
	openBrowser(authURL)

	select {
	case result := <-resultCh:
		return result, nil
	case err := <-errCh:
		return nil, err
	case <-ctx.Done():
		return nil, errfmt.New(errfmt.ExitAuth, "login timed out")
	}
}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		return
	}
	_ = cmd.Start()
}
