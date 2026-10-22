package api

import (
	"context"
	"crypto/tls"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/voska/qbo-cli/internal/errfmt"
)

const sentInvoice = `{"Invoice":{"Id":"5","DocNumber":"TEST-5","EmailStatus":"EmailSent","BillEmail":{"Address":"a+b@example.com"},"DeliveryInfo":{"DeliveryType":"Email","DeliveryTime":"2026-10-21T12:00:00Z"},"Line":[{"Description":"Fixture"}],"TotalAmt":25},"time":"fixture"}`

func TestSendInvoiceRequest(t *testing.T) {
	for _, to := range []string{"", "a+b@example.com", "a&b@example.com"} {
		for _, minor := range []int{0, 75} {
			t.Run(fmt.Sprintf("%s/minor=%d", to, minor), func(t *testing.T) {
				var calls atomic.Int32
				srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					calls.Add(1)
					if r.Method != http.MethodPost || r.URL.Path != "/v3/company/123/invoice/5/send" {
						t.Errorf("request = %s %s", r.Method, r.URL)
					}
					q := url.Values{}
					if to != "" {
						q.Set("sendTo", to)
					}
					if minor > 0 {
						q.Set("minorversion", fmt.Sprint(minor))
					}
					if !reflect.DeepEqual(r.URL.Query(), q) {
						t.Errorf("query = %q, want %v", r.URL.RawQuery, q)
					}
					if to != "" && !strings.Contains(r.URL.RawQuery, "sendTo="+url.QueryEscape(to)) {
						t.Errorf("unencoded recipient: %q", r.URL.RawQuery)
					}
					if r.Header.Get("Content-Type") != "application/octet-stream" || r.Header.Get("Accept") != "application/json" {
						t.Errorf("headers = %v", r.Header)
					}
					body, err := io.ReadAll(r.Body)
					if err != nil || len(body) != 0 {
						t.Errorf("body = %q, err = %v", body, err)
					}
					response := strings.ReplaceAll(sentInvoice, "a+b@example.com", to)
					if to == "" {
						response = sentInvoice
					}
					_, _ = io.WriteString(w, response)
				}))
				defer srv.Close()
				client := testClient(srv)
				client.minorVersion = minor
				got, err := client.SendInvoice(context.Background(), "5", to)
				if err != nil || calls.Load() != 1 {
					t.Fatalf("calls = %d, err = %v", calls.Load(), err)
				}
				invoice := got["Invoice"].(map[string]any)
				if invoice["DocNumber"] != "TEST-5" || invoice["TotalAmt"] != float64(25) || invoice["Line"] == nil || got["time"] != "fixture" {
					t.Errorf("response truncated: %v", got)
				}
			})
		}
	}
}

func TestSendInvoiceHTTPFailuresNeverRetry(t *testing.T) {
	for _, status := range []int{400, 401, 403, 404, 408, 429, 500, 502, 503, 504} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			var calls atomic.Int32
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				calls.Add(1)
				w.Header().Set("Retry-After", "0")
				w.WriteHeader(status)
				_, _ = io.WriteString(w, `{"Fault":{"type":"fixture"}}`)
			}))
			defer srv.Close()
			_, err := testClient(srv).SendInvoice(context.Background(), "5", "")
			want := errfmt.ExitError
			switch status {
			case 401:
				want = errfmt.ExitAuth
			case 403:
				want = errfmt.ExitForbidden
			case 404:
				want = errfmt.ExitNotFound
			}
			var e *errfmt.Error
			if !errors.As(err, &e) || e.Code != want || calls.Load() != 1 {
				t.Fatalf("calls = %d, err = %v, want code %d", calls.Load(), err, want)
			}
			if status == 408 || status == 429 || status >= 500 {
				assertUnknownSend(t, err)
			}
			if status == 400 && strings.Contains(err.Error(), "outcome unknown") {
				t.Errorf("validation failure mislabeled: %v", err)
			}
		})
	}
}

func assertUnknownSend(t *testing.T, err error) {
	t.Helper()
	var e *errfmt.Error
	if !errors.As(err, &e) || e.Code != errfmt.ExitError {
		t.Fatalf("err = %v, want code 1", err)
	}
	for _, text := range []string{"send outcome unknown", "invoice 5", "EmailStatus", "DeliveryInfo", "before sending again"} {
		if !strings.Contains(err.Error(), text) {
			t.Errorf("error missing %q: %v", text, err)
		}
	}
	if strings.Contains(err.Error(), "retry later") || strings.Contains(err.Error(), "safe to retry") {
		t.Errorf("unsafe retry hint: %v", err)
	}
}

func TestSendInvoiceUncertainResponse(t *testing.T) {
	for name, body := range map[string]string{
		"invalid JSON": "not json", "missing invoice": `{}`, "null": `null`,
		"wrong ID":          strings.Replace(sentInvoice, `"Id":"5"`, `"Id":"6"`, 1),
		"not sent":          strings.Replace(sentInvoice, "EmailSent", "NeedToSend", 1),
		"missing time":      strings.Replace(sentInvoice, "2026-10-21T12:00:00Z", "", 1),
		"wrong recipient":   strings.Replace(sentInvoice, "a+b@example.com", "other@example.com", 1),
		"missing recipient": strings.Replace(sentInvoice, "a+b@example.com", "", 1),
	} {
		t.Run(name, func(t *testing.T) {
			var calls atomic.Int32
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { calls.Add(1); _, _ = io.WriteString(w, body) }))
			defer srv.Close()
			_, err := testClient(srv).SendInvoice(context.Background(), "5", "a+b@example.com")
			assertUnknownSend(t, err)
			if calls.Load() != 1 {
				t.Fatalf("calls = %d", calls.Load())
			}
		})
	}
}

func TestSendInvoiceTransportFailure(t *testing.T) {
	for _, truncate := range []bool{false, true} {
		t.Run(fmt.Sprint(truncate), func(t *testing.T) {
			var calls atomic.Int32
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				calls.Add(1)
				conn, rw, err := w.(http.Hijacker).Hijack()
				if err != nil {
					t.Error(err)
					return
				}
				defer func() { _ = conn.Close() }()
				if truncate {
					_, _ = rw.WriteString("HTTP/1.1 200 OK\r\nContent-Length: 10000\r\n\r\n{")
					_ = rw.Flush()
				}
			}))
			defer srv.Close()
			_, err := testClient(srv).SendInvoice(context.Background(), "5", "")
			assertUnknownSend(t, err)
			if calls.Load() != 1 {
				t.Errorf("calls = %d", calls.Load())
			}
		})
	}
}

func TestSendInvoiceDoesNotFollowRedirects(t *testing.T) {
	for _, status := range []int{301, 302, 303, 307, 308} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			var calls atomic.Int32
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls.Add(1)
				http.Redirect(w, r, "/second-send", status)
			}))
			defer srv.Close()
			client := testClient(srv)
			_, err := client.SendInvoice(context.Background(), "5", "")
			assertUnknownSend(t, err)
			if calls.Load() != 1 {
				t.Errorf("calls = %d", calls.Load())
			}
			if client.httpClient.CheckRedirect != nil {
				t.Error("send changed the shared client's redirect policy")
			}
		})
	}
}

// Older supported Go HTTP/2 transports replay even a POST when its Body is nil
// after a peer's PROTOCOL_ERROR. A send must not take that retry path.
func TestSendInvoiceHTTP2ResetNeverReplays(t *testing.T) {
	certServer := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	certs := certServer.TLS.Certificates
	certServer.Close()
	listener, err := tls.Listen("tcp", "127.0.0.1:0", &tls.Config{Certificates: certs, NextProtos: []string{"h2"}})
	if err != nil {
		t.Fatal(err)
	}
	var calls atomic.Int32
	var workers sync.WaitGroup
	var mu sync.Mutex
	var connections []net.Conn
	workers.Add(1)
	go func() {
		defer workers.Done()
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			mu.Lock()
			connections = append(connections, conn)
			mu.Unlock()
			workers.Add(1)
			go func() {
				defer workers.Done()
				defer func() { _ = conn.Close() }()
				_ = conn.SetDeadline(time.Now().Add(5 * time.Second))
				preface := make([]byte, 24)
				if _, err := io.ReadFull(conn, preface); err != nil {
					return
				}
				if err := sendHTTP2Frame(conn, 4, 0, 0, nil); err != nil {
					return
				} // SETTINGS
				for {
					header := make([]byte, 9)
					if _, err := io.ReadFull(conn, header); err != nil {
						return
					}
					size := int(header[0])<<16 | int(header[1])<<8 | int(header[2])
					if _, err := io.CopyN(io.Discard, conn, int64(size)); err != nil {
						return
					}
					if header[3] == 4 && header[4]&1 == 0 { // SETTINGS ACK
						if err := sendHTTP2Frame(conn, 4, 1, 0, nil); err != nil {
							return
						}
					}
					if header[3] != 1 {
						continue
					} // HEADERS marks a dispatched request.
					code := byte(1) // PROTOCOL_ERROR, retryable in Go 1.25/1.26.
					if calls.Add(1) > 1 {
						code = 8
					} // CANCEL stops an accidental replay promptly.
					stream := binary.BigEndian.Uint32(header[5:]) & 0x7fffffff
					if err := sendHTTP2Frame(conn, 3, 0, stream, []byte{0, 0, 0, code}); err != nil {
						return
					}
				}
			}()
		}
	}()
	t.Cleanup(func() {
		_ = listener.Close()
		mu.Lock()
		for _, conn := range connections {
			_ = conn.Close()
		}
		mu.Unlock()
		workers.Wait()
	})
	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, ForceAttemptHTTP2: true} // Synthetic loopback TLS fixture.
	t.Cleanup(transport.CloseIdleConnections)
	client := &Client{httpClient: &http.Client{Transport: transport}, baseURL: "https://" + listener.Addr().String(), realmID: "123", minorVersion: 75}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = client.SendInvoice(ctx, "5", "")
	assertUnknownSend(t, err)
	if calls.Load() != 1 {
		t.Fatalf("send requests dispatched = %d, want exactly 1", calls.Load())
	}
}

func sendHTTP2Frame(w io.Writer, kind, flags byte, stream uint32, payload []byte) error {
	header := make([]byte, 9)
	header[0], header[1], header[2] = byte(len(payload)>>16), byte(len(payload)>>8), byte(len(payload))
	header[3], header[4] = kind, flags
	binary.BigEndian.PutUint32(header[5:], stream)
	_, err := w.Write(append(header, payload...))
	return err
}

func TestSendInvoiceHTTP2Success(t *testing.T) {
	var calls atomic.Int32
	srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		body, err := io.ReadAll(r.Body)
		if r.ProtoMajor != 2 || err != nil || len(body) != 0 {
			t.Errorf("protocol = %s, body = %q, err = %v", r.Proto, body, err)
		}
		_, _ = io.WriteString(w, sentInvoice)
	}))
	srv.EnableHTTP2 = true
	srv.StartTLS()
	defer srv.Close()
	result, err := testClient(srv).SendInvoice(context.Background(), "5", "a+b@example.com")
	if err != nil || result["Invoice"] == nil || calls.Load() != 1 {
		t.Fatalf("result = %v, calls = %d, err = %v", result, calls.Load(), err)
	}
}
