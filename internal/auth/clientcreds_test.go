package auth

import (
	"testing"

	"golang.org/x/oauth2"
)

// hermeticKeyring forces the encrypted-file backend in a temp dir so tests
// never touch the real OS keychain.
func hermeticKeyring(t *testing.T) {
	t.Helper()
	t.Setenv("QBO_KEYRING_BACKEND", "file")
	t.Setenv("QBO_CONFIG_DIR", t.TempDir())
}

func TestClientCredsRoundTrip(t *testing.T) {
	hermeticKeyring(t)

	if _, ok, err := LoadClientCreds(); err != nil || ok {
		t.Fatalf("expected no creds initially, got ok=%v err=%v", ok, err)
	}

	want := ClientCreds{
		ClientID:     "client-id",
		ClientSecret: "client-secret",
		RedirectURI:  "http://localhost:8844/callback",
	}
	if err := StoreClientCreds(want); err != nil {
		t.Fatalf("store: %v", err)
	}

	got, ok, err := LoadClientCreds()
	if err != nil || !ok {
		t.Fatalf("load after store: ok=%v err=%v", ok, err)
	}
	if got != want {
		t.Fatalf("round-trip mismatch: got %+v want %+v", got, want)
	}

	if err := DeleteClientCreds(); err != nil {
		t.Fatalf("delete: %v", err)
	}
	if _, ok, _ := LoadClientCreds(); ok {
		t.Fatalf("creds should be gone after delete")
	}
}

func TestDeleteClientCredsWhenAbsentIsNoError(t *testing.T) {
	hermeticKeyring(t)
	if err := DeleteClientCreds(); err != nil {
		t.Fatalf("delete of absent creds should not error: %v", err)
	}
}

func TestListTokenKeysExcludesClientCreds(t *testing.T) {
	hermeticKeyring(t)

	if err := StoreClientCreds(ClientCreds{ClientID: "id", ClientSecret: "s"}); err != nil {
		t.Fatal(err)
	}
	const realm = "9130353148918356"
	if err := StoreToken(realm, &oauth2.Token{AccessToken: "x"}); err != nil {
		t.Fatal(err)
	}

	keys, err := ListTokenKeys()
	if err != nil {
		t.Fatal(err)
	}

	var sawRealm bool
	for _, k := range keys {
		if k == clientCredsKey {
			t.Fatalf("ListTokenKeys leaked reserved client-creds key: %v", keys)
		}
		if k == realm {
			sawRealm = true
		}
	}
	if !sawRealm {
		t.Fatalf("expected company realm key in %v", keys)
	}
}
