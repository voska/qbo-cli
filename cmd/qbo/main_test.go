package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/voska/qbo-cli/internal/errfmt"
)

func TestSchemaUnknownCommandExit(t *testing.T) {
	if os.Getenv("QBO_TEST_SCHEMA_PROCESS") == "1" {
		os.Args = append([]string{"qbo"}, os.Args[3:]...)
		main()
		return
	}

	for _, flag := range []string{"", "--json", "--plain"} {
		t.Run("mode="+flag, func(t *testing.T) {
			args := []string{"-test.run=^TestSchemaUnknownCommandExit$", "--", "schema", "nosuch"}
			if flag != "" {
				args = append(args, flag)
			}
			process := exec.Command(os.Args[0], args...)
			process.Env = append(os.Environ(), "QBO_TEST_SCHEMA_PROCESS=1", "QBO_CONFIG_DIR="+t.TempDir(), "QBO_AUTO_JSON=0")
			var stdout, stderr bytes.Buffer
			process.Stdout, process.Stderr = &stdout, &stderr
			err := process.Run()
			var exitErr *exec.ExitError
			if !errors.As(err, &exitErr) || exitErr.ExitCode() != errfmt.ExitUsage {
				t.Fatalf("exit = %v, want %d; stdout = %q, stderr = %q", err, errfmt.ExitUsage, stdout.String(), stderr.String())
			}
			if stdout.Len() != 0 {
				t.Errorf("stdout = %q, want no success output", stdout.String())
			}
			if !strings.Contains(stderr.String(), "command not found: nosuch") {
				t.Errorf("stderr = %q, want unknown-command error", stderr.String())
			}
		})
	}
}
