package api

import "strings"

type EntityInfo struct {
	Name       string
	Endpoint   string
	Queryable  bool
	Creatable  bool
	Updatable  bool
	Deletable  bool
	ReadOnly   bool
}

var entities = map[string]EntityInfo{
	"account":         {Name: "Account", Endpoint: "account", Queryable: true, Creatable: true, Updatable: true},
	"bill":            {Name: "Bill", Endpoint: "bill", Queryable: true, Creatable: true, Updatable: true, Deletable: true},
	"billpayment":     {Name: "BillPayment", Endpoint: "billpayment", Queryable: true, Creatable: true, Updatable: true, Deletable: true},
	"budget":          {Name: "Budget", Endpoint: "budget", Queryable: true},
	"class":           {Name: "Class", Endpoint: "class", Queryable: true, Creatable: true, Updatable: true},
	"companyinfo":     {Name: "CompanyInfo", Endpoint: "companyinfo", Queryable: true, ReadOnly: true},
	"creditmemo":      {Name: "CreditMemo", Endpoint: "creditmemo", Queryable: true, Creatable: true, Updatable: true, Deletable: true},
	"customer":        {Name: "Customer", Endpoint: "customer", Queryable: true, Creatable: true, Updatable: true},
	"department":      {Name: "Department", Endpoint: "department", Queryable: true, Creatable: true, Updatable: true},
	"deposit":         {Name: "Deposit", Endpoint: "deposit", Queryable: true, Creatable: true, Updatable: true, Deletable: true},
	"employee":        {Name: "Employee", Endpoint: "employee", Queryable: true, Creatable: true, Updatable: true},
	"estimate":        {Name: "Estimate", Endpoint: "estimate", Queryable: true, Creatable: true, Updatable: true, Deletable: true},
	"invoice":         {Name: "Invoice", Endpoint: "invoice", Queryable: true, Creatable: true, Updatable: true, Deletable: true},
	"item":            {Name: "Item", Endpoint: "item", Queryable: true, Creatable: true, Updatable: true},
	"journalentry":    {Name: "JournalEntry", Endpoint: "journalentry", Queryable: true, Creatable: true, Updatable: true, Deletable: true},
	"payment":         {Name: "Payment", Endpoint: "payment", Queryable: true, Creatable: true, Updatable: true, Deletable: true},
	"paymentmethod":   {Name: "PaymentMethod", Endpoint: "paymentmethod", Queryable: true, Creatable: true, Updatable: true},
	"preferences":     {Name: "Preferences", Endpoint: "preferences", Queryable: true, Updatable: true},
	"purchase":        {Name: "Purchase", Endpoint: "purchase", Queryable: true, Creatable: true, Updatable: true, Deletable: true},
	"purchaseorder":   {Name: "PurchaseOrder", Endpoint: "purchaseorder", Queryable: true, Creatable: true, Updatable: true, Deletable: true},
	"refundreceipt":   {Name: "RefundReceipt", Endpoint: "refundreceipt", Queryable: true, Creatable: true, Updatable: true, Deletable: true},
	"salesreceipt":    {Name: "SalesReceipt", Endpoint: "salesreceipt", Queryable: true, Creatable: true, Updatable: true, Deletable: true},
	"taxcode":         {Name: "TaxCode", Endpoint: "taxcode", Queryable: true, ReadOnly: true},
	"taxrate":         {Name: "TaxRate", Endpoint: "taxrate", Queryable: true, ReadOnly: true},
	"term":            {Name: "Term", Endpoint: "term", Queryable: true, Creatable: true, Updatable: true},
	"timeactivity":    {Name: "TimeActivity", Endpoint: "timeactivity", Queryable: true, Creatable: true, Updatable: true, Deletable: true},
	"transfer":        {Name: "Transfer", Endpoint: "transfer", Queryable: true, Creatable: true, Updatable: true, Deletable: true},
	"vendor":          {Name: "Vendor", Endpoint: "vendor", Queryable: true, Creatable: true, Updatable: true},
	"vendorcredit":    {Name: "VendorCredit", Endpoint: "vendorcredit", Queryable: true, Creatable: true, Updatable: true, Deletable: true},
}

func LookupEntity(name string) (EntityInfo, bool) {
	key := strings.ToLower(strings.ReplaceAll(name, "-", ""))
	key = strings.TrimSuffix(key, "s")
	if e, ok := entities[key]; ok {
		return e, true
	}
	if e, ok := entities[name]; ok {
		return e, true
	}
	return EntityInfo{}, false
}

func AllEntities() map[string]EntityInfo {
	return entities
}
