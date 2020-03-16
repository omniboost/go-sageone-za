package sage_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetJournalEntries(t *testing.T) {
	req := client.NewGetJournalEntriesRequest()
	req.QueryParams().CompanyID = companyID
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
