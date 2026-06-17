package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// The delete flow must READ the current template, then POST the bare
// type-wrapper (Id + SyncToken + RecurringInfo) back with ?operation=delete.
func TestDeleteRecurringTransactionReadThenEcho(t *testing.T) {
	var postedBody map[string]any
	var deleteQuery string
	read := false

	mux := http.NewServeMux()
	srv := httptest.NewServer(mux)
	defer srv.Close()

	mux.HandleFunc("/v3/company/123/recurringtransaction/185", func(w http.ResponseWriter, _ *http.Request) {
		read = true
		_, _ = io.WriteString(w, `{"RecurringTransaction":{"Invoice":{"Id":"185","SyncToken":"3","RecurringInfo":{"Name":"X"}}},"time":"t"}`)
	})
	mux.HandleFunc("/v3/company/123/recurringtransaction", func(w http.ResponseWriter, r *http.Request) {
		deleteQuery = r.URL.RawQuery
		b, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(b, &postedBody)
		_, _ = io.WriteString(w, `{"RecurringTransaction":{"Invoice":{"status":"Deleted","Id":"185"}},"time":"t"}`)
	})

	resp, err := testClient(srv).DeleteRecurringTransaction(context.Background(), "185")
	if err != nil {
		t.Fatalf("DeleteRecurringTransaction: %v", err)
	}
	if !read {
		t.Error("delete did not read the current template first")
	}
	if deleteQuery != "operation=delete" {
		t.Errorf("delete query = %q, want operation=delete", deleteQuery)
	}
	inv, ok := postedBody["Invoice"].(map[string]any)
	if !ok {
		t.Fatalf("posted delete body should be the bare type wrapper, got %v", postedBody)
	}
	if inv["Id"] != "185" || inv["SyncToken"] != "3" {
		t.Errorf("Id/SyncToken not echoed back: %v", inv)
	}
	if _, ok := inv["RecurringInfo"]; !ok {
		t.Error("RecurringInfo must be echoed back for delete")
	}
	if _, ok := resp["RecurringTransaction"]; !ok {
		t.Errorf("delete response not returned: %v", resp)
	}
}

// recurringTxnBody must produce the bare type wrapper QBO's delete endpoint
// expects, regardless of how the read response wraps it, preserving Id,
// SyncToken, and RecurringInfo.
func TestRecurringTxnBody(t *testing.T) {
	tests := []struct {
		name string
		read string
	}{
		{
			name: "wrapped under RecurringTransaction",
			read: `{"RecurringTransaction":{"Invoice":{"Id":"185","SyncToken":"0","RecurringInfo":{"Name":"X"}}},"time":"t"}`,
		},
		{
			name: "bare type wrapper",
			read: `{"Invoice":{"Id":"185","SyncToken":"0","RecurringInfo":{"Name":"X"}},"time":"t"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var read map[string]any
			if err := json.Unmarshal([]byte(tt.read), &read); err != nil {
				t.Fatal(err)
			}
			body, err := recurringTxnBody(read)
			if err != nil {
				t.Fatalf("recurringTxnBody: %v", err)
			}
			var got map[string]any
			if err := json.Unmarshal(body, &got); err != nil {
				t.Fatal(err)
			}
			if _, ok := got["time"]; ok {
				t.Error("time envelope should be stripped")
			}
			if _, ok := got["RecurringTransaction"]; ok {
				t.Error("RecurringTransaction envelope should be unwrapped")
			}
			inv, ok := got["Invoice"].(map[string]any)
			if !ok {
				t.Fatalf("expected bare Invoice wrapper, got %v", got)
			}
			if inv["Id"] != "185" || inv["SyncToken"] != "0" {
				t.Errorf("Id/SyncToken not preserved: %v", inv)
			}
			if _, ok := inv["RecurringInfo"]; !ok {
				t.Error("RecurringInfo must be preserved for delete")
			}
		})
	}
}

func TestRecurringTxnBodyEmpty(t *testing.T) {
	if _, err := recurringTxnBody(map[string]any{"time": "t"}); err == nil {
		t.Error("expected an error when the read response has no body")
	}
}
