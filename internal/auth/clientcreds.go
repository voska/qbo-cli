package auth

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/99designs/keyring"
	"github.com/voska/qbo-cli/internal/errfmt"
)

// errNotFound reports a missing keyring entry across backends: the OS keychain
// returns keyring.ErrKeyNotFound, while the file backend surfaces a raw
// os.ErrNotExist.
func errNotFound(err error) bool {
	return errors.Is(err, keyring.ErrKeyNotFound) || errors.Is(err, os.ErrNotExist)
}

// clientCredsKey is the reserved keyring entry holding app-level OAuth client
// credentials. Realm IDs are numeric, so this hyphenated key never collides
// with a company token entry.
const clientCredsKey = "client-credentials"

// ClientCreds are the app-level OAuth credentials issued by the Intuit
// developer portal. They are shared across every connected company, so a
// single keyring entry holds them rather than one per realm.
type ClientCreds struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
}

func StoreClientCreds(c ClientCreds) error {
	kr, err := openKeyring()
	if err != nil {
		return errfmt.Wrap(errfmt.ExitConfig, "cannot open keyring", err)
	}
	data, err := json.Marshal(c)
	if err != nil {
		return errfmt.Wrap(errfmt.ExitError, "cannot marshal client credentials", err)
	}
	return kr.Set(keyring.Item{Key: clientCredsKey, Data: data})
}

// LoadClientCreds returns the stored credentials. The bool reports whether an
// entry existed; a missing entry is not an error so callers can fall back to
// the environment or config file.
func LoadClientCreds() (ClientCreds, bool, error) {
	kr, err := openKeyring()
	if err != nil {
		return ClientCreds{}, false, errfmt.Wrap(errfmt.ExitConfig, "cannot open keyring", err)
	}
	item, err := kr.Get(clientCredsKey)
	if errNotFound(err) {
		return ClientCreds{}, false, nil
	}
	if err != nil {
		return ClientCreds{}, false, errfmt.Wrap(errfmt.ExitConfig, "cannot read client credentials", err)
	}
	var c ClientCreds
	if err := json.Unmarshal(item.Data, &c); err != nil {
		return ClientCreds{}, false, errfmt.Wrap(errfmt.ExitConfig, "corrupt client credentials", err)
	}
	return c, true, nil
}

func DeleteClientCreds() error {
	kr, err := openKeyring()
	if err != nil {
		return errfmt.Wrap(errfmt.ExitConfig, "cannot open keyring", err)
	}
	if err := kr.Remove(clientCredsKey); err != nil && !errNotFound(err) {
		return errfmt.Wrap(errfmt.ExitConfig, "cannot remove client credentials", err)
	}
	return nil
}
