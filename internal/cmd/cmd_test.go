package cmd

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/voska/qbo-cli/internal/errfmt"
)

func exitCode(err error) int {
	var e *errfmt.Error
	if errors.As(err, &e) {
		return e.Code
	}
	return -1
}

func testGlobals(cli *CLI) *Globals {
	return &Globals{Ctx: context.Background(), CLI: cli}
}

func TestDeleteUnknownEntity(t *testing.T) {
	err := (&DeleteCmd{Entity: "bogus", ID: "1"}).Run(testGlobals(&CLI{}))
	if got := exitCode(err); got != errfmt.ExitUsage {
		t.Fatalf("exit = %d, want %d (%v)", got, errfmt.ExitUsage, err)
	}
}

// --no-input must refuse (fail), not silently delete — only --force authorizes it.
func TestDeleteNoInputWithoutForceFails(t *testing.T) {
	err := (&DeleteCmd{Entity: "invoice", ID: "5"}).Run(testGlobals(&CLI{NoInput: true}))
	if got := exitCode(err); got != errfmt.ExitUsage {
		t.Fatalf("--no-input delete should fail with ExitUsage, got exit %d (%v)", got, err)
	}
}

// Dry-run must return before constructing the API client / making any call.
func TestDeleteDryRunNoNetwork(t *testing.T) {
	if err := (&DeleteCmd{Entity: "invoice", ID: "5"}).Run(testGlobals(&CLI{DryRun: true})); err != nil {
		t.Fatalf("dry-run delete returned error: %v", err)
	}
}

func TestCreateUnknownEntity(t *testing.T) {
	err := (&CreateCmd{Entity: "bogus"}).Run(testGlobals(&CLI{}))
	if got := exitCode(err); got != errfmt.ExitUsage {
		t.Fatalf("exit = %d, want %d (%v)", got, errfmt.ExitUsage, err)
	}
}

func TestBatchRejectsEmptyAndInvalid(t *testing.T) {
	dir := t.TempDir()
	write := func(name, content string) string {
		p := filepath.Join(dir, name)
		if err := os.WriteFile(p, []byte(content), 0o600); err != nil {
			t.Fatal(err)
		}
		return p
	}

	empty := write("empty.json", "[]")
	if got := exitCode((&BatchCmd{File: empty}).Run(testGlobals(&CLI{}))); got != errfmt.ExitUsage {
		t.Errorf("empty batch: exit = %d, want %d", got, errfmt.ExitUsage)
	}

	bad := write("bad.json", "{not json")
	if got := exitCode((&BatchCmd{File: bad}).Run(testGlobals(&CLI{}))); got != errfmt.ExitUsage {
		t.Errorf("invalid batch: exit = %d, want %d", got, errfmt.ExitUsage)
	}
}
