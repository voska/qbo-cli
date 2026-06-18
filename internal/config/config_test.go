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
			if got := c.ResolveClientID(tc.keyring); got != tc.want {
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
			if got := c.ResolveClientSecret(tc.keyring); got != tc.want {
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
			if got := c.ResolveRedirectURI(tc.keyring); got != tc.want {
				t.Fatalf("got %q want %q", got, tc.want)
			}
		})
	}
}
