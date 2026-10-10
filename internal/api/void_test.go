package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"github.com/voska/qbo-cli/internal/errfmt"
)

func TestVoidReadThenPost(t *testing.T) {
	for _, name := range []string{"invoice", "payment", "salesreceipt", "billpayment"} {
		for _, minorVersion := range []int{0, 75} {
			t.Run(name+"/minor="+strconv.Itoa(minorVersion), func(t *testing.T) {
				entity, _ := LookupEntity(name)
				var methods []string
				wantResult := map[string]any{entity.Name: map[string]any{
					"Id": "5", "SyncToken": "8", "TotalAmt": float64(0), "PrivateNote": "Voided test", "DocNumber": "TEST-5",
				}, "time": "test"}
				srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					methods = append(methods, r.Method)
					if minorVersion > 0 && r.URL.Query().Get("minorversion") != strconv.Itoa(minorVersion) {
						t.Errorf("minorversion = %q", r.URL.Query().Get("minorversion"))
					}
					if r.Method == http.MethodGet {
						if r.URL.Path != "/v3/company/123/"+name+"/5" {
							t.Errorf("GET path = %q", r.URL.Path)
						}
						_ = json.NewEncoder(w).Encode(map[string]any{entity.Name: map[string]any{
							"Id": "5", "SyncToken": "7", "TotalAmt": 25,
						}})
						return
					}
					if r.URL.Path != "/v3/company/123/"+name {
						t.Errorf("POST path = %q", r.URL.Path)
					}
					wantOp, wantInclude := "void", ""
					wantBody := map[string]any{"Id": "5", "SyncToken": "7"}
					if name != "invoice" {
						wantOp, wantInclude = "update", "void"
						wantBody["sparse"] = true
					}
					if q := r.URL.Query(); q.Get("operation") != wantOp || q.Get("include") != wantInclude {
						t.Errorf("void query = %q", r.URL.RawQuery)
					}
					var body map[string]any
					if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
						t.Error(err)
					}
					if !reflect.DeepEqual(body, wantBody) {
						t.Errorf("void body = %v, want %v", body, wantBody)
					}
					_ = json.NewEncoder(w).Encode(wantResult)
				}))
				defer srv.Close()
				client := testClient(srv)
				client.minorVersion = minorVersion
				expectedToken := ""
				if minorVersion > 0 {
					expectedToken = "7"
				}
				result, alreadyVoid, err := client.Void(context.Background(), entity, "5", expectedToken)
				if err != nil {
					t.Fatal(err)
				}
				if alreadyVoid || !reflect.DeepEqual(methods, []string{"GET", "POST"}) {
					t.Errorf("alreadyVoid = %v, methods = %v", alreadyVoid, methods)
				}
				if !reflect.DeepEqual(result, wantResult) {
					t.Errorf("result = %v, want %v", result, wantResult)
				}
			})
		}
	}
}

func TestVoidReadGuards(t *testing.T) {
	tests := []struct {
		name        string
		read        string
		pin         string
		wantError   bool
		alreadyVoid bool
	}{
		{"already void", `{"Invoice":{"Id":"5","SyncToken":"7","TotalAmt":0,"PrivateNote":"Voided test"}}`, "", false, true},
		{"matching pin already void", `{"Invoice":{"Id":"5","SyncToken":"7","TotalAmt":0,"PrivateNote":"Voided"}}`, "7", false, true},
		{"stale pin", `{"Invoice":{"Id":"5","SyncToken":"7","TotalAmt":25}}`, "6", true, false},
		{"stale pin already void", `{"Invoice":{"Id":"5","SyncToken":"7","TotalAmt":0,"PrivateNote":"Voided"}}`, "6", true, false},
		{"missing wrapper", `{"time":"test"}`, "", true, false},
		{"missing token", `{"Invoice":{"Id":"5"}}`, "", true, false},
		{"empty token", `{"Invoice":{"Id":"5","SyncToken":""}}`, "", true, false},
		{"numeric token", `{"Invoice":{"Id":"5","SyncToken":7}}`, "", true, false},
		{"wrong ID", `{"Invoice":{"Id":"6","SyncToken":"7"}}`, "", true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var methods []string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				methods = append(methods, r.Method)
				_, _ = w.Write([]byte(tt.read))
			}))
			defer srv.Close()
			entity, _ := LookupEntity("invoice")
			result, alreadyVoid, err := testClient(srv).Void(context.Background(), entity, "5", tt.pin)
			if (err != nil) != tt.wantError || alreadyVoid != tt.alreadyVoid {
				t.Errorf("alreadyVoid = %v, err = %v", alreadyVoid, err)
			}
			if !reflect.DeepEqual(methods, []string{"GET"}) {
				t.Errorf("methods = %v, want GET only", methods)
			}
			if tt.alreadyVoid && result["Invoice"] == nil {
				t.Error("already-void response was discarded")
			}
		})
	}
}

func TestVoidRequiresBothMarkers(t *testing.T) {
	for _, read := range []string{
		`{"Invoice":{"Id":"5","SyncToken":"7","TotalAmt":0}}`,
		`{"Invoice":{"Id":"5","SyncToken":"7","TotalAmt":25,"PrivateNote":"Voided"}}`,
		`{"Invoice":{"Id":"5","SyncToken":"7","PrivateNote":"Voided"}}`,
	} {
		t.Run(read, func(t *testing.T) {
			posts := 0
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodGet {
					_, _ = w.Write([]byte(read))
					return
				}
				posts++
				_, _ = w.Write([]byte(`{"Invoice":{"Id":"5","SyncToken":"8","TotalAmt":0,"PrivateNote":"Voided"}}`))
			}))
			defer srv.Close()
			entity, _ := LookupEntity("invoice")
			_, alreadyVoid, err := testClient(srv).Void(context.Background(), entity, "5", "")
			if err != nil || alreadyVoid || posts != 1 {
				t.Errorf("posts = %d, alreadyVoid = %v, err = %v", posts, alreadyVoid, err)
			}
		})
	}
}

func TestVoidFailure(t *testing.T) {
	for _, tt := range []struct {
		name       string
		getStatus  int
		postStatus int
		post       string
		wantCode   int
		wantPosts  int
	}{
		{"read fails", 404, 200, `{}`, errfmt.ExitNotFound, 0},
		{"stale object at POST", 200, 400, `{"Fault":{"type":"ValidationFault"}}`, errfmt.ExitError, 1},
		{"unverified response", 200, 200, `{"Invoice":{"Id":"5","TotalAmt":25}}`, errfmt.ExitError, 1},
		{"wrong response ID", 200, 200, `{"Invoice":{"Id":"6","TotalAmt":0,"PrivateNote":"Voided"}}`, errfmt.ExitError, 1},
		{"missing response wrapper", 200, 200, `{}`, errfmt.ExitError, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			posts := 0
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodGet {
					w.WriteHeader(tt.getStatus)
					_, _ = w.Write([]byte(`{"Invoice":{"Id":"5","SyncToken":"7","TotalAmt":25}}`))
					return
				}
				posts++
				w.WriteHeader(tt.postStatus)
				_, _ = w.Write([]byte(tt.post))
			}))
			defer srv.Close()
			entity, _ := LookupEntity("invoice")
			_, _, err := testClient(srv).Void(context.Background(), entity, "5", "")
			var e *errfmt.Error
			if !errors.As(err, &e) || e.Code != tt.wantCode || posts != tt.wantPosts {
				t.Errorf("posts = %d, err = %v", posts, err)
			}
		})
	}
}
