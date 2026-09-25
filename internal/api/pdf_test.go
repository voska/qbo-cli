package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/voska/qbo-cli/internal/errfmt"
)

// The endpoint must be asked for application/pdf: QuickBooks rejects the
// application/json Accept header every other call sends.
func TestPDFRequestsPDF(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/company/123/invoice/7202/pdf" {
			t.Errorf("path = %s", r.URL.Path)
		}
		if got := r.Header.Get("Accept"); got != "application/pdf" {
			t.Errorf("Accept = %q, want application/pdf", got)
		}
		w.Header().Set("Content-Type", "application/pdf")
		_, _ = w.Write([]byte("%PDF-1.4 body"))
	}))
	defer srv.Close()

	got, err := testClient(srv).PDF(context.Background(), "invoice", "7202")
	if err != nil {
		t.Fatalf("PDF: %v", err)
	}
	if string(got) != "%PDF-1.4 body" {
		t.Fatalf("body = %q", got)
	}
}

// A 200 carrying a fault instead of a document must not be returned as a PDF.
func TestPDFRejectsNonPDFBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"Fault":{"Error":[{"Message":"boom"}]}}`))
	}))
	defer srv.Close()

	if _, err := testClient(srv).PDF(context.Background(), "invoice", "1"); err == nil {
		t.Fatal("expected an error for a non-PDF body")
	}
}

func TestPDFMapsHTTPErrors(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	_, err := testClient(srv).PDF(context.Background(), "invoice", "1")
	e, ok := err.(*errfmt.Error)
	if !ok || e.Code != errfmt.ExitNotFound {
		t.Fatalf("err = %v, want ExitNotFound", err)
	}
}

func TestSupportsPDF(t *testing.T) {
	if !SupportsPDF("invoice") || !SupportsPDF("estimate") {
		t.Fatal("invoice and estimate must render as PDF")
	}
	if SupportsPDF("customer") {
		t.Fatal("customer has no PDF rendering")
	}
}
