package cmd

import (
	"reflect"
	"testing"

	"github.com/alecthomas/kong"
)

func TestSchemaIncludesRegisteredCommands(t *testing.T) {
	parser, err := kong.New(&CLI{})
	if err != nil {
		t.Fatal(err)
	}
	commands, ok := fullSchema("test")["commands"].([]map[string]any)
	if !ok {
		t.Fatal("schema has no command list")
	}
	listed := make(map[string]bool, len(commands))
	for _, command := range commands {
		name, ok := command["name"].(string)
		if !ok || name == "" {
			t.Fatalf("schema command has no name: %v", command)
		}
		listed[name] = true
	}
	for _, command := range parser.Model.Children {
		if command.Type != kong.CommandNode {
			continue
		}
		t.Run(command.Name, func(t *testing.T) {
			if !listed[command.Name] {
				t.Errorf("registered command %q is missing from schema", command.Name)
			}
		})
	}
}

func TestPDFSchemaMatchesCommand(t *testing.T) {
	var cli CLI
	parser, err := kong.New(&cli)
	if err != nil {
		t.Fatal(err)
	}
	var command *kong.Node
	for _, node := range parser.Model.Children {
		if node.Type == kong.CommandNode && node.Name == "pdf" {
			command = node
			break
		}
	}
	if command == nil {
		t.Fatal("pdf is missing from the command model")
	}
	schema, err := commandSchema("pdf", "test")
	if err != nil {
		t.Fatal(err)
	}
	if schema["name"] != command.Name || schema["help"] != command.Help {
		t.Fatalf("pdf schema does not match the command: %v", schema)
	}
	if !reflect.DeepEqual(schema["args"], []string{"entity", "id"}) {
		t.Errorf("pdf args = %v, want entity and id", schema["args"])
	}
	if !reflect.DeepEqual(schema["flags"], []string{"--output"}) {
		t.Errorf("pdf flags = %v, want --output", schema["flags"])
	}
}
