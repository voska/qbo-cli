package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/alecthomas/kong"
	"github.com/voska/qbo-cli/internal/errfmt"
)

const sendResponse = `{"Invoice":{"Id":"5","DocNumber":"TEST-5","EmailStatus":"EmailSent","BillEmail":{"Address":"a+b@example.com"},"DeliveryInfo":{"DeliveryType":"Email","DeliveryTime":"2026-10-21T12:00:00Z"},"Line":[{"Description":"Fixture"}],"TotalAmt":25},"time":"fixture"}`

func TestSendCommandAvailable(t *testing.T) {
	var cli CLI
	parser, err := kong.New(&cli)
	if err != nil {
		t.Fatal(err)
	}
	ctx, err := parser.Parse([]string{"send", "invoice", "5", "--to", "a+b@example.com", "--dry-run", "--no-input"})
	if err != nil {
		t.Fatal(err)
	}
	if err := ctx.Run(testGlobals(&cli)); err != nil {
		t.Fatal(err)
	}
}

func TestSendDryRunNoAuthOrNetwork(t *testing.T) {
	requests := voidCommandServer(t, func(_ http.ResponseWriter, _ *http.Request) { t.Error("unexpected request") })
	for _, to := range []string{"", "a+b@example.com"} {
		command := &SendCmd{Entity: "invoice", ID: "5"}
		if to != "" {
			command.To = &to
		}
		stdout, stderr, err := captureVoidOutput(t, func() error { return command.Run(testGlobals(&CLI{DryRun: true, NoInput: true})) })
		if err != nil || stdout != "" || requests.Load() != 0 {
			t.Fatalf("stdout = %q, requests = %d, err = %v", stdout, requests.Load(), err)
		}
		for _, text := range []string{"POST /v3/company/{id}/invoice/5/send", "third party", "Cc/Bcc"} {
			if !strings.Contains(stderr, text) {
				t.Errorf("dry run missing %q: %s", text, stderr)
			}
		}
		if to != "" {
			if !strings.Contains(stderr, "sendTo=a%2Bb%40example.com") || !strings.Contains(stderr, "replaces") || !strings.Contains(stderr, "BillEmail.Address") {
				t.Errorf("dry run = %q", stderr)
			}
		} else if strings.Contains(stderr, "sendTo=") || !strings.Contains(stderr, "invoice's BillEmail") {
			t.Errorf("dry run = %q", stderr)
		}
	}
}

func TestSendRejectsInputWithoutAuthOrNetwork(t *testing.T) {
	requests := voidCommandServer(t, func(_ http.ResponseWriter, _ *http.Request) { t.Error("unexpected request") })
	for _, entity := range []string{"bogus", "customer", "estimate", "payment", "salesreceipt", "creditmemo"} {
		if err := (&SendCmd{Entity: entity, ID: "5"}).Run(testGlobals(&CLI{Force: true})); exitCode(err) != errfmt.ExitUsage {
			t.Errorf("entity %q: %v", entity, err)
		}
	}
	for _, id := range []string{"", "INV-5", "5/../6", "5?sendTo=x@example.com", "5#", "%35", "5\n"} {
		if err := (&SendCmd{Entity: "invoice", ID: id}).Run(testGlobals(&CLI{Force: true})); exitCode(err) != errfmt.ExitUsage {
			t.Errorf("id %q: %v", id, err)
		}
	}
	for _, to := range []string{"", "bad", "a@example.com,b@example.com", "Name <a@example.com>", " a@example.com", "a@example.com\r\nBcc:b@example.com"} {
		if err := (&SendCmd{Entity: "invoice", ID: "5", To: &to}).Run(testGlobals(&CLI{Force: true})); exitCode(err) != errfmt.ExitUsage {
			t.Errorf("to %q: %v", to, err)
		}
	}
	for _, noInput := range []bool{false, true} {
		withSendStdin(t, "y\n", func() {
			err := (&SendCmd{Entity: "invoice", ID: "5"}).Run(testGlobals(&CLI{NoInput: noInput}))
			if exitCode(err) != errfmt.ExitUsage || !strings.Contains(err.Error(), "--force") {
				t.Errorf("no-input=%v: %v", noInput, err)
			}
		})
	}
	if requests.Load() != 0 {
		t.Errorf("requests = %d", requests.Load())
	}
}

func withSendStdin(t *testing.T, text string, run func()) {
	t.Helper()
	file, err := os.CreateTemp(t.TempDir(), "stdin")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = file.Close() }()
	if _, err = file.WriteString(text); err != nil {
		t.Fatal(err)
	}
	if _, err = file.Seek(0, 0); err != nil {
		t.Fatal(err)
	}
	original := os.Stdin
	os.Stdin = file
	defer func() { os.Stdin = original }()
	run()
}

func TestSendCommandOutput(t *testing.T) {
	t.Setenv("QBO_AUTO_JSON", "0")
	for _, mode := range []string{"human", "plain", "json", "select"} {
		t.Run(mode, func(t *testing.T) {
			requests := voidCommandServer(t, func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" || r.URL.Path != "/v3/company/123/invoice/5/send" || r.URL.Query().Get("sendTo") != "a+b@example.com" {
					t.Errorf("request = %s %s", r.Method, r.URL)
				}
				_, _ = io.WriteString(w, sendResponse)
			})
			cli := &CLI{Force: true, NoInput: true, JSON: mode == "json" || mode == "select", Plain: mode == "plain", CompanyID: "123", MinorVer: 75}
			if mode == "select" {
				cli.Select = []string{"Invoice.Id", "Invoice.EmailStatus"}
			}
			g, err := NewGlobals(testGlobals(cli).Ctx, cli)
			if err != nil {
				t.Fatal(err)
			}
			to := "a+b@example.com"
			stdout, stderr, err := captureVoidOutput(t, func() error { return (&SendCmd{Entity: "invoices", ID: "5", To: &to}).Run(g) })
			if err != nil || requests.Load() != 1 {
				t.Fatalf("requests = %d, err = %v", requests.Load(), err)
			}
			if strings.Contains(stderr, "[y/N]") {
				t.Errorf("force prompted: %q", stderr)
			}
			switch mode {
			case "json":
				var got, want any
				if err := json.Unmarshal([]byte(stdout), &got); err != nil {
					t.Fatal(err)
				}
				_ = json.Unmarshal([]byte(sendResponse), &want)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("JSON truncated: %s", stdout)
				}
			case "select":
				var got map[string]any
				if err := json.Unmarshal([]byte(stdout), &got); err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(got, map[string]any{"Invoice.Id": "5", "Invoice.EmailStatus": "EmailSent"}) {
					t.Errorf("selection = %v", got)
				}
			default:
				for _, text := range []string{"Id", "DocNumber", "EmailStatus", "BillEmail.Address", "DeliveryInfo.DeliveryTime", "TEST-5", "EmailSent", "a+b@example.com", "2026-10-21T12:00:00Z"} {
					if !strings.Contains(stdout, text) {
						t.Errorf("output missing %q: %s", text, stdout)
					}
				}
				if mode == "plain" && !strings.Contains(stdout, "\t") {
					t.Errorf("not TSV: %s", stdout)
				}
			}
		})
	}
}

func TestSendCommandFailureCodes(t *testing.T) {
	for status, want := range map[int]int{400: 1, 401: 4, 403: 6, 404: 5, 429: 1, 500: 1} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			requests := voidCommandServer(t, func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(status) })
			cli := &CLI{Force: true, NoInput: true, CompanyID: "123", MinorVer: 75}
			g, err := NewGlobals(testGlobals(cli).Ctx, cli)
			if err != nil {
				t.Fatal(err)
			}
			err = (&SendCmd{Entity: "invoice", ID: "5"}).Run(g)
			if exitCode(err) != want || requests.Load() != 1 {
				t.Errorf("requests = %d, err = %v", requests.Load(), err)
			}
		})
	}
}

func TestSendSchema(t *testing.T) {
	schema, err := commandSchema("send", "test")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(schema["args"], []string{"entity", "id"}) || !reflect.DeepEqual(schema["flags"], []string{"--to"}) {
		t.Errorf("send schema = %v", schema)
	}
	if !reflect.DeepEqual(schema["entities"], []string{"invoice"}) || schema["retryable"] != false {
		t.Errorf("send safety schema = %v", schema)
	}
	for _, text := range []string{"third party", "BillEmail.Address", "never retried"} {
		if !strings.Contains(schema["help"].(string), text) {
			t.Errorf("schema help missing %q: %v", text, schema)
		}
	}
}

func TestSendConfirmation(t *testing.T) {
	for _, recipient := range []string{"the invoice's BillEmail", "a+b@example.com (replaces the invoice's BillEmail.Address)"} {
		for _, answer := range []string{"", "n\n", "yes\n", "y extra\n", "y\n", "Y\n"} {
			t.Run(recipient+"/"+answer, func(t *testing.T) {
				withSendStdin(t, answer, func() {
					stdout, stderr, err := captureVoidOutput(t, func() error {
						if got := confirmSend("5", recipient); got != (answer == "y\n" || answer == "Y\n") {
							t.Errorf("confirmation for %q = %v", answer, got)
						}
						return nil
					})
					if err != nil || stdout != "" {
						t.Fatalf("stdout=%q, err=%v", stdout, err)
					}
					for _, text := range []string{"invoice 5", recipient, "third party", "[y/N]"} {
						if !strings.Contains(stderr, text) {
							t.Errorf("prompt missing %q: %s", text, stderr)
						}
					}
				})
			})
		}
	}
}
