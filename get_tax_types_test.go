package sage_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetTaxTypes(t *testing.T) {
	req := client.NewGetTaxTypesRequest()
	req.QueryParams().CompanyID = companyID
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
