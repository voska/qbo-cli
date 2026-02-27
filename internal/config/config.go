package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/voska/qbo-cli/internal/errfmt"
)

const (
	EnvProduction = "production"
	EnvSandbox    = "sandbox"
)

type Company struct {
	RealmID     string `json:"realm_id"`
	CompanyName string `json:"company_name,omitempty"`
	Environment string `json:"environment"`
}

type Config struct {
	DefaultCompany string    `json:"default_company,omitempty"`
	Companies      []Company `json:"companies,omitempty"`
	ClientID       string    `json:"client_id,omitempty"`
	ClientSecret   string    `json:"client_secret,omitempty"`
}

func Dir() (string, error) {
	if d := os.Getenv("QBO_CONFIG_DIR"); d != "" {
		return d, nil
	}
	home, err := os.UserConfigDir()
	if err != nil {
		return "", errfmt.Wrap(errfmt.ExitConfig, "cannot determine config directory", err)
	}
	return filepath.Join(home, "qbo"), nil
}

func Load() (*Config, error) {
	dir, err := Dir()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(dir, "config.json")
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &Config{}, nil
	}
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitConfig, "cannot read config", err)
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, errfmt.Wrap(errfmt.ExitConfig, "invalid config", err)
	}
	return &cfg, nil
}

func (c *Config) Save() error {
	dir, err := Dir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return errfmt.Wrap(errfmt.ExitConfig, "cannot create config directory", err)
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return errfmt.Wrap(errfmt.ExitConfig, "cannot marshal config", err)
	}
	return os.WriteFile(filepath.Join(dir, "config.json"), data, 0o600)
}

func (c *Config) FindCompany(realmID string) *Company {
	for i := range c.Companies {
		if c.Companies[i].RealmID == realmID {
			return &c.Companies[i]
		}
	}
	return nil
}

func (c *Config) AddOrUpdateCompany(co Company) {
	for i := range c.Companies {
		if c.Companies[i].RealmID == co.RealmID {
			c.Companies[i] = co
			return
		}
	}
	c.Companies = append(c.Companies, co)
}

func (c *Config) ResolveCompanyID(flagValue string) string {
	if flagValue != "" {
		return flagValue
	}
	if v := os.Getenv("QBO_COMPANY_ID"); v != "" {
		return v
	}
	return c.DefaultCompany
}

func (c *Config) ResolveClientID() string {
	if v := os.Getenv("QBO_CLIENT_ID"); v != "" {
		return v
	}
	return c.ClientID
}

func (c *Config) ResolveClientSecret() string {
	if v := os.Getenv("QBO_CLIENT_SECRET"); v != "" {
		return v
	}
	return c.ClientSecret
}
