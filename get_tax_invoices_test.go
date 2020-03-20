package sage_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetTaxInvoices(t *testing.T) {
	req := client.NewGetTaxInvoicesRequest()
	req.QueryParams().IncludeDetail = true
	req.QueryParams().CompanyID = companyID
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
