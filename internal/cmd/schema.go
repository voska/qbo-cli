package cmd

import (
	"sort"

	"github.com/voska/qbo-cli/internal/api"
)

type SchemaCmd struct {
	Command string `arg:"" optional:"" help:"Show schema for a specific command."`
}

func (c *SchemaCmd) Run(g *Globals) error {
	if c.Command != "" {
		return WriteOutput(g.Ctx, commandSchema(c.Command, g.Version))
	}
	return WriteOutput(g.Ctx, fullSchema(g.Version))
}

func fullSchema(version string) map[string]any {
	entities := api.AllEntities()
	entityNames := make([]string, 0, len(entities))
	for k := range entities {
		entityNames = append(entityNames, k)
	}
	sort.Strings(entityNames)

	entityList := make([]map[string]any, 0, len(entities))
	for _, name := range entityNames {
		e := entities[name]
		entityList = append(entityList, map[string]any{
			"name":      e.Name,
			"key":       name,
			"queryable": e.Queryable,
			"creatable": e.Creatable,
			"updatable": e.Updatable,
			"deletable": e.Deletable,
			"read_only": e.ReadOnly,
		})
	}

	return map[string]any{
		"name":    "qbo",
		"version": version,
		"commands": []map[string]any{
			{
				"name": "auth",
				"help": "Authentication (login, logout, status, refresh)",
				"subcommands": []map[string]any{
					{"name": "login", "help": "Authenticate with QuickBooks Online", "flags": []string{"--manual"}},
					{"name": "logout", "help": "Remove stored credentials"},
					{"name": "status", "help": "Show current auth status"},
					{"name": "refresh", "help": "Force token refresh"},
				},
			},
			{
				"name":  "list",
				"help":  "List entities (sugar for query)",
				"args":  []string{"entity"},
				"flags": []string{"--where", "--order-by", "--limit", "--offset"},
			},
			{
				"name": "get",
				"help": "Read a single entity by ID",
				"args": []string{"entity", "id"},
			},
			{
				"name":  "create",
				"help":  "Create an entity from JSON",
				"args":  []string{"entity"},
				"flags": []string{"--file"},
			},
			{
				"name":  "update",
				"help":  "Update an entity",
				"args":  []string{"entity", "id"},
				"flags": []string{"--file", "--sparse"},
			},
			{
				"name": "delete",
				"help": "Delete an entity by ID",
				"args": []string{"entity", "id"},
			},
			{
				"name": "query",
				"help": "Run a raw query against the QBO API",
				"args": []string{"query"},
			},
			{
				"name":  "batch",
				"help":  "Run batch operations",
				"flags": []string{"--file"},
			},
			{
				"name":  "cdc",
				"help":  "Change data capture polling",
				"flags": []string{"--entities", "--since"},
			},
			{
				"name":  "report",
				"help":  "Run a QBO report",
				"args":  []string{"type"},
				"flags": []string{"--start-date", "--end-date"},
			},
			{
				"name": "company",
				"help": "Company info and switching",
				"subcommands": []map[string]any{
					{"name": "info", "help": "Show company info from QBO API"},
					{"name": "list", "help": "List configured companies"},
					{"name": "switch", "help": "Set default company", "args": []string{"realm-id"}},
				},
			},
			{
				"name": "schema",
				"help": "Dump CLI schema as JSON for agent introspection",
				"args": []string{"command?"},
			},
			{
				"name": "exit-codes",
				"help": "Print exit code table",
			},
		},
		"global_flags": []map[string]any{
			{"name": "--json", "short": "-j", "help": "Output JSON to stdout", "aliases": []string{"--machine"}},
			{"name": "--plain", "short": "-p", "help": "Output TSV, no color", "aliases": []string{"--tsv"}},
			{"name": "--select", "help": "Project output to fields", "aliases": []string{"--fields"}},
			{"name": "--results-only", "help": "Strip pagination metadata"},
			{"name": "--company-id", "help": "Target company realm ID", "env": "QBO_COMPANY_ID"},
			{"name": "--sandbox", "help": "Use sandbox environment"},
			{"name": "--dry-run", "short": "-n", "help": "Show what would happen", "aliases": []string{"--noop", "--preview"}},
			{"name": "--no-input", "help": "Never prompt"},
			{"name": "--force", "help": "Skip confirmations", "aliases": []string{"--yes", "--assume-yes"}},
			{"name": "--verbose", "short": "-v", "help": "Verbose output to stderr"},
			{"name": "--minor-version", "help": "QBO API minor version", "default": "75"},
		},
		"entities": entityList,
	}
}

func commandSchema(name, version string) map[string]any {
	schema := fullSchema(version)
	commands, ok := schema["commands"].([]map[string]any)
	if !ok {
		return map[string]any{"error": "command not found: " + name}
	}
	for _, cmd := range commands {
		if cmd["name"] == name {
			return cmd
		}
	}
	return map[string]any{"error": "command not found: " + name}
}
