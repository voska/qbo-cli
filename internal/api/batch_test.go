package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Batch must pass the parsed items through under BatchItemRequest (regression
// test for the bug where the command discarded its input and posted nil).
func TestBatchPostsItems(t *testing.T) {
	var gotBody map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3/company/123/batch" {
			t.Errorf("path = %q", r.URL.Path)
		}
		b, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(b, &gotBody)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"BatchItemResponse":[{"bId":"1"}]}`))
	}))
	defer srv.Close()

	items := []json.RawMessage{
		json.RawMessage(`{"bId":"1","operation":"create","Invoice":{"x":1}}`),
		json.RawMessage(`{"bId":"2","Query":"select * from Invoice"}`),
	}
	resp, err := testClient(srv).Batch(context.Background(), items)
	if err != nil {
		t.Fatalf("Batch: %v", err)
	}

	arr, ok := gotBody["BatchItemRequest"].([]any)
	if !ok || len(arr) != 2 {
		t.Fatalf("BatchItemRequest not passed through: %v", gotBody)
	}
	first, _ := arr[0].(map[string]any)
	if first["bId"] != "1" {
		t.Errorf("first item not preserved: %v", first)
	}
	if _, ok := resp["BatchItemResponse"]; !ok {
		t.Errorf("response not returned: %v", resp)
	}
}
