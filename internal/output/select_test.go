package output

import (
	"encoding/json"
	"testing"
)

func TestProjectFields(t *testing.T) {
	data := map[string]any{
		"Id":          "123",
		"DisplayName": "Acme Corp",
		"Balance":     100.0,
		"BillAddr": map[string]any{
			"City":  "San Jose",
			"State": "CA",
		},
	}

	result := ProjectFields(data, []string{"Id", "DisplayName", "BillAddr.City"})
	m, ok := result.(map[string]any)
	if !ok {
		t.Fatal("expected map[string]any")
	}
	if m["Id"] != "123" {
		t.Errorf("Id = %v, want 123", m["Id"])
	}
	if m["DisplayName"] != "Acme Corp" {
		t.Errorf("DisplayName = %v, want Acme Corp", m["DisplayName"])
	}
	if m["BillAddr.City"] != "San Jose" {
		t.Errorf("BillAddr.City = %v, want San Jose", m["BillAddr.City"])
	}
	if _, ok := m["Balance"]; ok {
		t.Error("Balance should not be present")
	}
}

func TestProjectFieldsSlice(t *testing.T) {
	data := []any{
		map[string]any{"Id": "1", "Name": "A"},
		map[string]any{"Id": "2", "Name": "B"},
	}

	result := ProjectFields(data, []string{"Id"})
	arr, ok := result.([]any)
	if !ok {
		t.Fatal("expected []any")
	}
	if len(arr) != 2 {
		t.Fatalf("len = %d, want 2", len(arr))
	}
	m := arr[0].(map[string]any)
	if m["Id"] != "1" {
		t.Errorf("Id = %v, want 1", m["Id"])
	}
	if _, ok := m["Name"]; ok {
		t.Error("Name should not be present")
	}
}

func TestStripMetadata(t *testing.T) {
	raw := `{"QueryResponse":{"Invoice":[{"Id":"1"},{"Id":"2"}],"startPosition":1,"maxResults":2,"totalCount":2}}`
	var data any
	_ = json.Unmarshal([]byte(raw), &data)

	result := StripMetadata(data)
	arr, ok := result.([]any)
	if !ok {
		t.Fatal("expected []any")
	}
	if len(arr) != 2 {
		t.Fatalf("len = %d, want 2", len(arr))
	}
}
