package config

import "testing"

func TestResolveClientIDPrecedence(t *testing.T) {
	c := &Config{ClientID: "from-config"}
	cases := []struct {
		name, env, keyring, want string
	}{
		{"env wins over all", "from-env", "from-keyring", "from-env"},
		{"keyring over config", "", "from-keyring", "from-keyring"},
		{"config fallback", "", "", "from-config"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("QBO_CLIENT_ID", tc.env)
			if got := c.ResolveClientID(func() string { return tc.keyring }); got != tc.want {
				t.Fatalf("got %q want %q", got, tc.want)
			}
		})
	}
}

func TestResolveClientSecretPrecedence(t *testing.T) {
	c := &Config{ClientSecret: "from-config"}
	cases := []struct {
		name, env, keyring, want string
	}{
		{"env wins over all", "from-env", "from-keyring", "from-env"},
		{"keyring over config", "", "from-keyring", "from-keyring"},
		{"config fallback", "", "", "from-config"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("QBO_CLIENT_SECRET", tc.env)
			if got := c.ResolveClientSecret(func() string { return tc.keyring }); got != tc.want {
				t.Fatalf("got %q want %q", got, tc.want)
			}
		})
	}
}

func TestResolveRedirectURIPrecedence(t *testing.T) {
	c := &Config{RedirectURI: "from-config"}
	cases := []struct {
		name, env, keyring, want string
	}{
		{"env wins over all", "from-env", "from-keyring", "from-env"},
		{"keyring over config", "", "from-keyring", "from-keyring"},
		{"config fallback", "", "", "from-config"},
		{"empty when unset", "", "", "from-config"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("QBO_REDIRECT_URI", tc.env)
			if got := c.ResolveRedirectURI(func() string { return tc.keyring }); got != tc.want {
				t.Fatalf("got %q want %q", got, tc.want)
			}
		})
	}
}

// When the env var is set it must win without ever invoking the keyring
// provider — so an env-only setup never opens the keychain.
func TestResolversSkipKeyringWhenEnvSet(t *testing.T) {
	c := &Config{ClientID: "cfg-id", ClientSecret: "cfg-secret", RedirectURI: "cfg-uri"}
	t.Setenv("QBO_CLIENT_ID", "env-id")
	t.Setenv("QBO_CLIENT_SECRET", "env-secret")
	t.Setenv("QBO_REDIRECT_URI", "env-uri")

	mustNotCall := func() string {
		t.Helper()
		t.Fatal("keyring provider must not be consulted when env var is set")
		return ""
	}

	if got := c.ResolveClientID(mustNotCall); got != "env-id" {
		t.Errorf("ClientID = %q, want env-id", got)
	}
	if got := c.ResolveClientSecret(mustNotCall); got != "env-secret" {
		t.Errorf("ClientSecret = %q, want env-secret", got)
	}
	if got := c.ResolveRedirectURI(mustNotCall); got != "env-uri" {
		t.Errorf("RedirectURI = %q, want env-uri", got)
	}
}
