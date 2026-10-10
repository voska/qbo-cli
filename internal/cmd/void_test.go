package cmd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/alecthomas/kong"
	"github.com/voska/qbo-cli/internal/auth"
	"github.com/voska/qbo-cli/internal/errfmt"
	"golang.org/x/oauth2"
)

func TestVoidCommandAvailable(t *testing.T) {
	var cli CLI
	parser, err := kong.New(&cli)
	if err != nil {
		t.Fatal(err)
	}
	ctx, err := parser.Parse([]string{"void", "invoice", "5", "--dry-run", "--no-input"})
	if err != nil {
		t.Fatalf("void command is unavailable: %v", err)
	}
	if err := ctx.Run(testGlobals(&cli)); err != nil {
		t.Fatal(err)
	}
}

type voidRoundTripper func(*http.Request) (*http.Response, error)

func (f voidRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// Redirect the real command's OAuth transport to a fixture, using only a fake
// token in an isolated file keyring. No production credentials are consulted.
func voidCommandServer(t *testing.T, handler http.HandlerFunc) *atomic.Int32 {
	t.Helper()
	t.Setenv("QBO_CONFIG_DIR", t.TempDir())
	t.Setenv("QBO_KEYRING_BACKEND", "file")
	if err := auth.StoreToken("123", &oauth2.Token{AccessToken: "fixture", Expiry: time.Now().Add(time.Hour)}); err != nil {
		t.Fatal(err)
	}
	var requests atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests.Add(1)
		handler(w, r)
	}))
	t.Cleanup(srv.Close)
	fixtureURL, err := url.Parse(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	original := http.DefaultTransport
	http.DefaultTransport = voidRoundTripper(func(r *http.Request) (*http.Response, error) {
		r = r.Clone(r.Context())
		r.URL.Scheme, r.URL.Host = fixtureURL.Scheme, fixtureURL.Host
		return original.RoundTrip(r)
	})
	t.Cleanup(func() { http.DefaultTransport = original })
	return &requests
}

func captureVoidOutput(t *testing.T, run func() error) (string, string, error) {
	t.Helper()
	out, err := os.CreateTemp(t.TempDir(), "stdout")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = out.Close() }()
	hints, err := os.CreateTemp(t.TempDir(), "stderr")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = hints.Close() }()
	stdout, stderr := os.Stdout, os.Stderr
	os.Stdout, os.Stderr = out, hints
	defer func() { os.Stdout, os.Stderr = stdout, stderr }()
	runErr := run()
	outBytes, err := os.ReadFile(out.Name())
	if err != nil {
		t.Fatal(err)
	}
	hintBytes, err := os.ReadFile(hints.Name())
	if err != nil {
		t.Fatal(err)
	}
	return string(outBytes), string(hintBytes), runErr
}

func TestVoidDryRunNoNetwork(t *testing.T) {
	requests := voidCommandServer(t, func(_ http.ResponseWriter, _ *http.Request) { t.Error("unexpected HTTP request") })
	for _, name := range []string{"invoice", "payment", "salesreceipt", "billpayment"} {
		for _, token := range []string{"", "7"} {
			t.Run(name+"/token="+token, func(t *testing.T) {
				stdout, stderr, err := captureVoidOutput(t, func() error {
					return (&VoidCmd{Entity: name, ID: "5", SyncToken: token}).Run(testGlobals(&CLI{DryRun: true, NoInput: true}))
				})
				if err != nil || stdout != "" || requests.Load() != 0 {
					t.Fatalf("stdout = %q, requests = %d, err = %v", stdout, requests.Load(), err)
				}
				query := "operation=void"
				if name != "invoice" {
					query = "operation=update&include=void"
				}
				if !strings.Contains(stderr, "GET then POST /v3/company/{id}/"+name+"?"+query) {
					t.Errorf("dry-run plan = %q", stderr)
				}
				if token == "" {
					if !strings.Contains(stderr, "SyncToken read at execution") {
						t.Errorf("dry-run invents a token: %q", stderr)
					}
					return
				}
				bodyStart := strings.Index(stderr, "{\"Id\"")
				if bodyStart < 0 {
					t.Fatalf("dry-run body is missing: %q", stderr)
				}
				bodyLine := strings.TrimSpace(stderr[bodyStart:])
				var body map[string]any
				if err := json.Unmarshal([]byte(bodyLine), &body); err != nil {
					t.Fatalf("dry-run body is not JSON: %q (%v)", stderr, err)
				}
				if body["Id"] != "5" || body["SyncToken"] != "7" || (name != "invoice" && body["sparse"] != true) {
					t.Errorf("dry-run body = %v", body)
				}
			})
		}
	}
}

func TestVoidUsageNoNetwork(t *testing.T) {
	requests := voidCommandServer(t, func(_ http.ResponseWriter, _ *http.Request) { t.Error("unexpected HTTP request") })
	for _, name := range []string{"bogus", "customer", "bill", "recurringtransaction"} {
		err := (&VoidCmd{Entity: name, ID: "5"}).Run(testGlobals(&CLI{Force: true}))
		if exitCode(err) != errfmt.ExitUsage || (name == "customer" && err.Error() != "Customer cannot be voided") {
			t.Errorf("%s: %v", name, err)
		}
	}
	err := (&VoidCmd{Entity: "invoice", ID: "5"}).Run(testGlobals(&CLI{NoInput: true}))
	if exitCode(err) != errfmt.ExitUsage || requests.Load() != 0 {
		t.Errorf("requests = %d, err = %v", requests.Load(), err)
	}
}

func TestVoidConfirmationCancel(t *testing.T) {
	stdin, err := os.CreateTemp(t.TempDir(), "stdin")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = stdin.Close() }()
	original := os.Stdin
	os.Stdin = stdin // EOF defaults to cancellation.
	defer func() { os.Stdin = original }()
	stdout, stderr, err := captureVoidOutput(t, func() error {
		return (&VoidCmd{Entity: "invoice", ID: "5"}).Run(testGlobals(&CLI{}))
	})
	if err != nil || stdout != "" || !strings.Contains(stderr, "can't be undone") || !strings.Contains(stderr, "cancelled") {
		t.Errorf("stdout = %q, stderr = %q, err = %v", stdout, stderr, err)
	}
}

func TestVoidCommandOutput(t *testing.T) {
	for _, alreadyVoid := range []bool{false, true} {
		t.Run(map[bool]string{false: "void", true: "already void"}[alreadyVoid], func(t *testing.T) {
			response := `{"Invoice":{"Id":"5","SyncToken":"8","TotalAmt":0,"PrivateNote":"Voided test","DocNumber":"TEST-5"},"time":"test"}`
			requests := voidCommandServer(t, func(w http.ResponseWriter, r *http.Request) {
				if r.Method == "GET" && !alreadyVoid {
					_, _ = w.Write([]byte(`{"Invoice":{"Id":"5","SyncToken":"7","TotalAmt":25}}`))
					return
				}
				_, _ = w.Write([]byte(response))
			})
			cli := &CLI{JSON: true, Force: true, NoInput: true, CompanyID: "123", MinorVer: 75}
			g, err := NewGlobals(testGlobals(cli).Ctx, cli)
			if err != nil {
				t.Fatal(err)
			}
			stdout, stderr, err := captureVoidOutput(t, func() error {
				return (&VoidCmd{Entity: "invoices", ID: "5"}).Run(g)
			})
			if err != nil {
				t.Fatal(err)
			}
			var got map[string]any
			if err := json.Unmarshal([]byte(stdout), &got); err != nil {
				t.Fatalf("stdout = %q (%v)", stdout, err)
			}
			invoice := got["Invoice"].(map[string]any)
			if invoice["TotalAmt"] != float64(0) || invoice["PrivateNote"] != "Voided test" || invoice["DocNumber"] != "TEST-5" || got["time"] != "test" {
				t.Errorf("response was changed: %v", got)
			}
			wantRequests := int32(2)
			if alreadyVoid {
				wantRequests = 1
			}
			if requests.Load() != wantRequests || strings.Contains(stderr, "already voided") != alreadyVoid {
				t.Errorf("requests = %d, stderr = %q", requests.Load(), stderr)
			}
		})
	}
}

func TestVoidSchema(t *testing.T) {
	schema, err := commandSchema("void", "test")
	if err != nil {
		t.Fatal(err)
	}
	if schema["name"] != "void" {
		t.Fatalf("void is missing from schema: %v", schema)
	}
	flags, ok := schema["flags"].([]string)
	if !ok || len(flags) != 1 || flags[0] != "--sync-token" {
		t.Fatalf("void flags = %v, want --sync-token", schema["flags"])
	}
	entities := fullSchema("test")["entities"].([]map[string]any)
	for _, entity := range entities {
		key := entity["key"]
		want := key == "invoice" || key == "payment" || key == "salesreceipt" || key == "billpayment"
		if entity["voidable"] != want {
			t.Errorf("%v voidable = %v, want %v", key, entity["voidable"], want)
		}
	}
}
