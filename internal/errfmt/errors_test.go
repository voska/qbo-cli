package errfmt

import (
	"errors"
	"testing"
)

func TestErrorMessage(t *testing.T) {
	e := New(ExitAuth, "not authenticated")
	if e.Error() != "not authenticated" {
		t.Errorf("got %q", e.Error())
	}
	if e.Code != ExitAuth {
		t.Errorf("code = %d, want %d", e.Code, ExitAuth)
	}
}

func TestErrorWithDetail(t *testing.T) {
	e := Wrap(ExitError, "request failed", errors.New("connection refused"))
	want := "request failed: connection refused"
	if e.Error() != want {
		t.Errorf("got %q, want %q", e.Error(), want)
	}
}

func TestExitCodeTable(t *testing.T) {
	table := ExitCodeTable()
	if len(table) != 10 {
		t.Errorf("len = %d, want 10", len(table))
	}
	if table[0].Code != ExitOK {
		t.Errorf("first code = %d, want %d", table[0].Code, ExitOK)
	}
}
