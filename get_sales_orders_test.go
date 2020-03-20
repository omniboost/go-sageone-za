package sage_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetSalesOrders(t *testing.T) {
	req := client.NewGetSalesOrdersRequest()
	req.QueryParams().CompanyID = companyID
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
