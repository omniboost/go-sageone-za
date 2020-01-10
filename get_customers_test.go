package aktiva_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetCustomers(t *testing.T) {
	req := client.NewGetCustomersRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
