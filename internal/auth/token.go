package auth

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/99designs/keyring"
	"github.com/voska/qbo-cli/internal/errfmt"
	"golang.org/x/oauth2"
)

const keyringService = "qbo-cli"

func tokenDir() string {
	if d := os.Getenv("QBO_CONFIG_DIR"); d != "" {
		return filepath.Join(d, "tokens")
	}
	home, err := os.UserConfigDir()
	if err != nil {
		home = os.TempDir()
	}
	return filepath.Join(home, "qbo", "tokens")
}

func openKeyring() (keyring.Keyring, error) {
	dir := tokenDir()
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, errfmt.Wrap(errfmt.ExitConfig, "cannot create token directory", err)
	}
	return keyring.Open(keyring.Config{
		ServiceName:                    keyringService,
		KeychainTrustApplication:       true,
		KeychainAccessibleWhenUnlocked: true,
		FileDir:                        dir,
		FilePasswordFunc:               func(_ string) (string, error) { return "", nil },
	})
}

func StoreToken(realmID string, token *oauth2.Token) error {
	kr, err := openKeyring()
	if err != nil {
		return errfmt.Wrap(errfmt.ExitConfig, "cannot open keyring", err)
	}
	data, err := json.Marshal(token)
	if err != nil {
		return errfmt.Wrap(errfmt.ExitError, "cannot marshal token", err)
	}
	return kr.Set(keyring.Item{
		Key:  realmID,
		Data: data,
	})
}

func LoadToken(realmID string) (*oauth2.Token, error) {
	kr, err := openKeyring()
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitConfig, "cannot open keyring", err)
	}
	item, err := kr.Get(realmID)
	if err != nil {
		return nil, errfmt.Auth("not authenticated — run: qbo auth login")
	}
	var token oauth2.Token
	if err := json.Unmarshal(item.Data, &token); err != nil {
		return nil, errfmt.Wrap(errfmt.ExitConfig, "corrupt token data", err)
	}
	return &token, nil
}

func DeleteToken(realmID string) error {
	kr, err := openKeyring()
	if err != nil {
		return errfmt.Wrap(errfmt.ExitConfig, "cannot open keyring", err)
	}
	return kr.Remove(realmID)
}

func ListTokenKeys() ([]string, error) {
	kr, err := openKeyring()
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitConfig, "cannot open keyring", err)
	}
	keys, err := kr.Keys()
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func IsTokenExpired(token *oauth2.Token) bool {
	return token.Expiry.Before(time.Now())
}
