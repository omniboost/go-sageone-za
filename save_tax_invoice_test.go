package sage_test

import (
	"encoding/json"
	"log"
	"testing"

	sage "github.com/omniboost/go-sageone-za"
)

func TestSaveTaxInvoice(t *testing.T) {
	body := sage.SaveTaxInvoiceRequestBody{}

	req := client.NewSaveTaxInvoiceRequest()
	req.SetRequestBody(body)
	req.QueryParams().CompanyID = companyID
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
