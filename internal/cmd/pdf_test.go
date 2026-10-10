package cmd

import (
	"testing"

	"github.com/voska/qbo-cli/internal/errfmt"
)

func TestPDFRejectsBadInput(t *testing.T) {
	cases := []struct{ name, entity, id string }{
		{"unknown entity", "bogus", "1"},
		{"entity without a PDF rendering", "customer", "1"},
		{"document number instead of Id", "invoice", "INV-1467"},
		{"path smuggled into the id", "invoice", "1/../2"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := (&PDFCmd{Entity: tc.entity, ID: tc.id}).Run(testGlobals(&CLI{}))
			if got := exitCode(err); got != errfmt.ExitUsage {
				t.Fatalf("exit = %d, want %d (%v)", got, errfmt.ExitUsage, err)
			}
		})
	}
}

// Dry-run must return before constructing the API client / making any call.
func TestPDFDryRunNoNetwork(t *testing.T) {
	if err := (&PDFCmd{Entity: "invoice", ID: "5"}).Run(testGlobals(&CLI{DryRun: true})); err != nil {
		t.Fatalf("dry-run pdf returned error: %v", err)
	}
}
