package sage_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetCompanies(t *testing.T) {
	req := client.NewGetCompaniesRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
