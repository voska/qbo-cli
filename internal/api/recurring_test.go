package api

import (
	"encoding/json"
	"testing"
)

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
