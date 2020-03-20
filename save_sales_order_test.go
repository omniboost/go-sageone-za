package sage_test

import (
	"encoding/json"
	"log"
	"testing"

	sage "github.com/omniboost/go-sageone-za"
)

func TestSaveSalesOrder(t *testing.T) {
	body := sage.SaveSalesOrderRequestBody{}

	req := client.NewSaveSalesOrderRequest()
	req.SetRequestBody(body)
	req.QueryParams().CompanyID = companyID
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
