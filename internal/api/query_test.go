package api

import "testing"

func TestQueryBuilder(t *testing.T) {
	tests := []struct {
		name string
		fn   func() string
		want string
	}{
		{
			name: "simple select all",
			fn:   func() string { return NewQuery("Invoice").Build() },
			want: "SELECT * FROM Invoice MAXRESULTS 1000",
		},
		{
			name: "with where",
			fn: func() string {
				return NewQuery("Invoice").Where("Balance > '0'").Build()
			},
			want: "SELECT * FROM Invoice WHERE Balance > '0' MAXRESULTS 1000",
		},
		{
			name: "with order by",
			fn: func() string {
				return NewQuery("Invoice").OrderBy("DueDate DESC").Build()
			},
			want: "SELECT * FROM Invoice ORDERBY DueDate DESC MAXRESULTS 1000",
		},
		{
			name: "with pagination",
			fn: func() string {
				return NewQuery("Customer").MaxResults(10).StartPosition(11).Build()
			},
			want: "SELECT * FROM Customer STARTPOSITION 11 MAXRESULTS 10",
		},
		{
			name: "full query",
			fn: func() string {
				return NewQuery("Invoice").
					Select("Id, DocNumber, Balance").
					Where("Balance > '0'").
					OrderBy("DueDate").
					MaxResults(50).
					StartPosition(1).
					Build()
			},
			want: "SELECT Id, DocNumber, Balance FROM Invoice WHERE Balance > '0' ORDERBY DueDate MAXRESULTS 50",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestLookupEntity(t *testing.T) {
	tests := []struct {
		input    string
		wantName string
		wantOK   bool
	}{
		{"invoice", "Invoice", true},
		{"invoices", "Invoice", true},
		{"Invoice", "Invoice", true},
		{"customer", "Customer", true},
		{"customers", "Customer", true},
		{"bill-payment", "BillPayment", true},
		{"unknown", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			e, ok := LookupEntity(tt.input)
			if ok != tt.wantOK {
				t.Errorf("LookupEntity(%q) ok = %v, want %v", tt.input, ok, tt.wantOK)
			}
			if ok && e.Name != tt.wantName {
				t.Errorf("LookupEntity(%q) name = %q, want %q", tt.input, e.Name, tt.wantName)
			}
		})
	}
}
